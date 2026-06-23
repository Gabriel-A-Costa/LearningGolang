package handler

import (
	"encoding/json"
	"errors"
	"net/http"

	"go.uber.org/zap"

	"github.com/gabriel-a-costa/LearningGolang/jun/zap/api/internal/domain"
	"github.com/gabriel-a-costa/LearningGolang/jun/zap/api/internal/service"
)

// ProductHandler é a camada HTTP. Tem duas responsabilidades:
//  1. Traduzir erros de domínio em respostas HTTP
//  2. Logar erros inesperados (5xx)
type ProductHandler struct {
	svc    *service.ProductService
	logger *zap.Logger
}

func New(svc *service.ProductService, logger *zap.Logger) *ProductHandler {
	return &ProductHandler{
		svc: svc,
		// Named logger: todo log daqui terá {"logger":"product.handler"}.
		// Facilita filtrar logs em produção: grep 'product.handler'.
		logger: logger.Named("product.handler"),
	}
}

// ══════════════════════════════════════════════════════════════════════════════
// REGRAS DA CAMADA HANDLER
//
// RESPONSABILIDADE: traduzir erros em HTTP e logar o que importa.
//
// AQUI é onde erros são TRATADOS (não apenas propagados).
// Para cada erro, o handler toma 3 decisões:
//   1. Qual HTTP status retornar?
//   2. Qual mensagem mostrar ao cliente? (nunca detalhes internos!)
//   3. Logar? Em qual nível?
//
// ÁRVORE DE DECISÃO DE LOG:
//
//   ErrProductNotFound  → 404 → zap.Debug  (esperado, não alarma ninguém)
//   ErrProductDuplicate → 409 → zap.Debug  (esperado, fluxo normal)
//   ErrInvalidProduct   → 400 → zap.Debug  (erro do cliente, não nosso)
//   qualquer outro      → 500 → zap.Error  (NOSSO problema, ops precisa saber)
//
// POR QUE não logar 4xx como Error?
//   zap.Error geralmente dispara alertas (PagerDuty, Datadog).
//   404 e 400 são da responsabilidade do CLIENTE, não do servidor.
//   Logar todo 404 como Error é o anti-padrão "menino que gritou lobo":
//   quando um erro real acontecer, estará enterrado no ruído.
//
// COMO errors.Is() funciona aqui:
//   O erro recebido pode ser:
//   "service.GetByID: repository.FindByID id=X: product not found"
//   errors.Is(err, domain.ErrProductNotFound) → true
//   Porque errors.Is() DESEMBRULHA a cadeia até encontrar uma correspondência.
//   Isso só funciona se TODAS as camadas usaram %w (não %v).
//
// COMO errors.As() funciona aqui:
//   var ve *domain.ValidationError
//   errors.As(err, &ve) → true quando existe um *ValidationError na cadeia
//   Diferente de errors.Is() que checa VALOR, errors.As() extrai o TIPO.
//   Use errors.As() quando precisa acessar campos da struct (Field, Message).
// ══════════════════════════════════════════════════════════════════════════════

// Create — POST /products
func (h *ProductHandler) Create(w http.ResponseWriter, r *http.Request) {
	log := h.logger.With(
		zap.String("method", r.Method),
		zap.String("path", r.URL.Path),
	)

	var input domain.CreateInput
	if err := json.NewDecoder(r.Body).Decode(&input); err != nil {
		// JSON malformado é erro do cliente. Debug, não Error.
		log.Debug("corpo da requisição inválido", zap.Error(err))
		writeError(w, http.StatusBadRequest, "corpo JSON inválido")
		return
	}

	product, err := h.svc.Create(input)
	if err != nil {
		h.handleError(w, log, err)
		return
	}

	writeJSON(w, http.StatusCreated, product)
}

// GetByID — GET /products/{id}
func (h *ProductHandler) GetByID(w http.ResponseWriter, r *http.Request) {
	id := r.PathValue("id")
	log := h.logger.With(
		zap.String("method", r.Method),
		zap.String("path", r.URL.Path),
		zap.String("product_id", id),
	)

	product, err := h.svc.GetByID(id)
	if err != nil {
		h.handleError(w, log, err)
		return
	}

	writeJSON(w, http.StatusOK, product)
}

// List — GET /products
func (h *ProductHandler) List(w http.ResponseWriter, r *http.Request) {
	log := h.logger.With(
		zap.String("method", r.Method),
		zap.String("path", r.URL.Path),
	)

	products, err := h.svc.List()
	if err != nil {
		h.handleError(w, log, err)
		return
	}

	writeJSON(w, http.StatusOK, products)
}

// Update — PUT /products/{id}
func (h *ProductHandler) Update(w http.ResponseWriter, r *http.Request) {
	id := r.PathValue("id")
	log := h.logger.With(
		zap.String("method", r.Method),
		zap.String("path", r.URL.Path),
		zap.String("product_id", id),
	)

	var input domain.UpdateInput
	if err := json.NewDecoder(r.Body).Decode(&input); err != nil {
		log.Debug("corpo da requisição inválido", zap.Error(err))
		writeError(w, http.StatusBadRequest, "corpo JSON inválido")
		return
	}

	product, err := h.svc.Update(id, input)
	if err != nil {
		h.handleError(w, log, err)
		return
	}

	writeJSON(w, http.StatusOK, product)
}

// Delete — DELETE /products/{id}
func (h *ProductHandler) Delete(w http.ResponseWriter, r *http.Request) {
	id := r.PathValue("id")
	log := h.logger.With(
		zap.String("method", r.Method),
		zap.String("path", r.URL.Path),
		zap.String("product_id", id),
	)

	if err := h.svc.Delete(id); err != nil {
		h.handleError(w, log, err)
		return
	}

	w.WriteHeader(http.StatusNoContent)
}

// handleError é o PONTO CENTRAL de roteamento de erros.
//
// Centralizar aqui garante consistência: nenhum handler individual
// decide por conta própria o que logar. Uma regra, aplicada em todo lugar.
//
// Esta função demonstra os 3 cenários que você pediu:
func (h *ProductHandler) handleError(w http.ResponseWriter, log *zap.Logger, err error) {
	switch {

	// ── Cenário 1: Produto não encontrado ───────────────────────────────────
	//
	// errors.Is() percorre:
	//   "service.GetByID: repository.FindByID id=X: product not found"
	//                                                 ↑
	//                                         encontra ErrProductNotFound aqui
	//
	// Retorno: 404
	// Log: Debug (esperado, não significa problema)
	// NÃO incluímos zap.Error(err) porque a mensagem do log já é suficiente.
	case errors.Is(err, domain.ErrProductNotFound):
		log.Debug("produto não encontrado")
		writeError(w, http.StatusNotFound, "produto não encontrado")

	case errors.Is(err, domain.ErrProductDuplicate):
		log.Debug("produto já existe")
		writeError(w, http.StatusConflict, "produto já existe")

	// ── Cenário 3: Dados inválidos enviados pelo cliente ────────────────────
	//
	// Aqui demonstramos errors.Is() E errors.As() ao mesmo tempo:
	//
	//   errors.Is(err, domain.ErrInvalidProduct) → verifica SE é erro de validação
	//   errors.As(err, &ve) → extrai o *ValidationError para acessar Field e Message
	//
	// Retorno: 400 com detalhe do campo que falhou
	// Log: Debug (o cliente enviou dado errado, não é problema nosso)
	case errors.Is(err, domain.ErrInvalidProduct):
		var ve *domain.ValidationError
		if errors.As(err, &ve) {
			// errors.As extraiu o tipo concreto — podemos acessar Field e Message.
			// Isso dá ao cliente uma resposta muito mais útil: {"field":"name","message":"é obrigatório"}
			log.Debug("requisição inválida", zap.String("field", ve.Field), zap.String("message", ve.Message))
			writeJSON(w, http.StatusBadRequest, map[string]string{
				"error":   "dados inválidos",
				"field":   ve.Field,
				"message": ve.Message,
			})
		} else {
			log.Debug("validação falhou", zap.Error(err))
			writeError(w, http.StatusBadRequest, err.Error())
		}

	// ── Cenário 2: Erro inesperado de infraestrutura ─────────────────────────
	//
	// Tudo que não é erro de domínio conhecido cai aqui.
	// Este é O ÚNICO lugar onde usamos zap.Error() para problemas reais.
	//
	// zap.Error(err) imprime a cadeia COMPLETA:
	//   "error": "service.GetByID: repository.FindByID id=X: database connection failed"
	//
	// Para o cliente: mensagem GENÉRICA. Nunca exponha detalhes internos!
	// Detalhes internos (nomes de tabelas, IPs, etc.) ajudam atacantes.
	//
	// Retorno: 500
	// Log: Error → dispara alertas em produção (PagerDuty, Datadog)
	default:
		log.Error("erro interno inesperado", zap.Error(err))
		writeError(w, http.StatusInternalServerError, "erro interno do servidor")
	}
}

// ─── Helpers HTTP ────────────────────────────────────────────────────────────

func writeJSON(w http.ResponseWriter, status int, body any) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	json.NewEncoder(w).Encode(body) //nolint:errcheck
}

func writeError(w http.ResponseWriter, status int, msg string) {
	writeJSON(w, status, map[string]string{"error": msg})
}
