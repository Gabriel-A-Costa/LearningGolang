package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"time"

	"go.uber.org/zap"

	"github.com/gabriel-a-costa/LearningGolang/jun/zap/api/internal/handler"
	"github.com/gabriel-a-costa/LearningGolang/jun/zap/api/internal/repository"
	"github.com/gabriel-a-costa/LearningGolang/jun/zap/api/internal/service"
	zaplogger "github.com/gabriel-a-costa/LearningGolang/jun/zap/api/pkg/logger"
)

func main() {
	// Logger é o primeiro recurso inicializado.
	// Qualquer erro de startup também deve ser estruturado.
	logger, err := zaplogger.New()
	if err != nil {
		panic(fmt.Sprintf("falha ao inicializar logger: %v", err))
	}
	defer logger.Sync() // garante flush do buffer ao encerrar

	// ── Injeção de Dependência (manual, sem framework) ──────────────────────
	// Construímos de baixo para cima: Repository → Service → Handler
	// Cada camada só conhece a camada imediatamente abaixo.
	repo := repository.New()
	svc := service.New(repo)
	h := handler.New(svc, logger)

	// ── Rotas (Go 1.22+ routing nativo com {id}) ────────────────────────────
	mux := http.NewServeMux()
	mux.HandleFunc("POST /products", h.Create)
	mux.HandleFunc("GET /products", h.List)
	mux.HandleFunc("GET /products/{id}", h.GetByID)
	mux.HandleFunc("PUT /products/{id}", h.Update)
	mux.HandleFunc("DELETE /products/{id}", h.Delete)

	// ── Endpoints de debug (APENAS para o demo educacional) ─────────────────
	// Ativa/desativa a simulação de falha de banco para testar o Cenário 2.
	mux.HandleFunc("POST /debug/db-error", func(w http.ResponseWriter, r *http.Request) {
		repo.SimulateDBError(true)
		logger.Warn("simulação de erro de banco ATIVADA")
		respondJSON(w, http.StatusOK, map[string]string{"status": "db error simulation ON"})
	})
	mux.HandleFunc("DELETE /debug/db-error", func(w http.ResponseWriter, r *http.Request) {
		repo.SimulateDBError(false)
		logger.Warn("simulação de erro de banco DESATIVADA")
		respondJSON(w, http.StatusOK, map[string]string{"status": "db error simulation OFF"})
	})

	server := &http.Server{
		Addr:         ":8085",
		Handler:      mux,
		ReadTimeout:  5 * time.Second,
		WriteTimeout: 10 * time.Second,
		IdleTimeout:  60 * time.Second,
	}

	logger.Info("servidor iniciado", zap.String("addr", server.Addr))
	printInstructions()

	if err := server.ListenAndServe(); err != nil && err != http.ErrServerClosed {
		logger.Fatal("servidor falhou", zap.Error(err))
	}
}

func respondJSON(w http.ResponseWriter, status int, body any) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	json.NewEncoder(w).Encode(body) //nolint:errcheck
}

func printInstructions() {
	fmt.Print(`
╔══════════════════════════════════════════════════════════════════════════════╗
║                    PRODUCT API — Demo de Error Handling                     ║
╠══════════════════════════════════════════════════════════════════════════════╣
║                                                                              ║
║  CENÁRIO 1 — Produto não encontrado (404, log.Debug, sem alarme)             ║
║    curl -s localhost:8080/products/id-inexistente | jq                       ║
║                                                                              ║
║  CENÁRIO 2 — Erro de banco (500, log.Error, dispara alerta)                  ║
║    curl -X POST localhost:8080/debug/db-error                                ║
║    curl -s localhost:8080/products | jq        <- verá o erro 500            ║
║    curl -X DELETE localhost:8080/debug/db-error                              ║
║                                                                              ║
║  CENÁRIO 3 — Validação inválida (400, log.Debug, sem alarme)                 ║
║    curl -s -X POST localhost:8080/products \                                  ║
║      -H 'Content-Type: application/json' \                                   ║
║      -d '{"name":"","price":-5}' | jq                                        ║
║                                                                              ║
║  HAPPY PATH — Criar produto                                                  ║
║    curl -s -X POST localhost:8080/products \                                  ║
║      -H 'Content-Type: application/json' \                                   ║
║      -d '{"name":"Teclado Mecanico","price":299.90,"stock":10}' | jq         ║
║                                                                              ║
╚══════════════════════════════════════════════════════════════════════════════╝
`)
}
