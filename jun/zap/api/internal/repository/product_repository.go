package repository

import (
	"errors"
	"fmt"
	"sync"

	"github.com/gabriel-a-costa/LearningGolang/jun/zap/api/internal/domain"
)

// ErrDBConnection simula falhas de infraestrutura (rede, timeout, etc.).
// Fica no package repository porque é uma preocupação de infraestrutura,
// não de negócio. O handler trata qualquer erro que não seja sentinel de
// domínio como erro interno (500).
var ErrDBConnection = errors.New("database connection failed")

// InMemoryProductRepository é uma implementação thread-safe em memória.
// Em produção, substituiria por PostgresProductRepository, etc.
// A interface domain.Repository garantiria que ambas são intercambiáveis.
type InMemoryProductRepository struct {
	mu              sync.RWMutex
	products        map[string]domain.Product
	simulateDBError bool // apenas para o demo educacional
}

func New() *InMemoryProductRepository {
	return &InMemoryProductRepository{
		products: make(map[string]domain.Product),
	}
}

// SimulateDBError ativa/desativa a simulação de falha de banco.
// Exposto apenas para o demo — NUNCA faça isso em produção.
func (r *InMemoryProductRepository) SimulateDBError(enabled bool) {
	r.simulateDBError = enabled
}

// ══════════════════════════════════════════════════════════════════════════════
// REGRAS DA CAMADA REPOSITORY
//
// RESPONSABILIDADE: acesso a dados. Nada mais.
//
// O QUE RETORNAR:
//   - Erros sentinel de domínio para erros de negócio (ErrProductNotFound, etc.)
//   - Erros de infraestrutura para falhas técnicas (ErrDBConnection, etc.)
//   - SEMPRE embrulhe com fmt.Errorf("repository.Operacao: %w", err)
//
// O QUE NÃO LOGAR:
//   - ABSOLUTAMENTE NADA. Zero logs aqui.
//
// POR QUE NÃO LOGAR?
//   O repository não sabe o que significa este erro para o chamador.
//   Um "product not found" durante um GET é NORMAL (retorna 404).
//   O mesmo erro durante uma lookup obrigatória pode ser CRÍTICO.
//   Somente o handler (camada superior) tem contexto suficiente para decidir.
//
//   Se você logar aqui E no handler → log duplicado para o mesmo erro.
//   Isso polui seus logs e confunde quem está debugando.
//
// POR QUE %w E NÃO %v?
//   fmt.Errorf("repo: %v", err) → QUEBRA errors.Is() — não use
//   fmt.Errorf("repo: %w", err) → PRESERVA a cadeia — sempre use
// ══════════════════════════════════════════════════════════════════════════════

func (r *InMemoryProductRepository) Create(product domain.Product) error {
	// Cenário 2: erro de infraestrutura
	if r.simulateDBError {
		return fmt.Errorf("repository.Create: %w", ErrDBConnection)
	}

	r.mu.Lock()
	defer r.mu.Unlock()

	if _, exists := r.products[product.ID]; exists {
		// Erro de negócio: produto já existe.
		// %w preserva ErrProductDuplicate para errors.Is() upstream.
		return fmt.Errorf("repository.Create id=%s: %w", product.ID, domain.ErrProductDuplicate)
	}

	r.products[product.ID] = product
	return nil
}

// FindByID demonstra os 3 cenários de erro em uma única função.
func (r *InMemoryProductRepository) FindByID(id string) (domain.Product, error) {
	// Cenário 2: o banco está inacessível
	if r.simulateDBError {
		// ErrDBConnection NÃO é um sentinel de domínio.
		// O handler vai cair no caso "default" e logar como Error (500).
		return domain.Product{}, fmt.Errorf("repository.FindByID id=%s: %w", id, ErrDBConnection)
	}

	r.mu.RLock()
	defer r.mu.RUnlock()

	p, ok := r.products[id]
	if !ok {
		// Cenário 1: produto não encontrado (erro de negócio esperado).
		// O handler vai reconhecer ErrProductNotFound e retornar 404.
		// Nenhum log aqui — o handler decide se isso merece atenção.
		return domain.Product{}, fmt.Errorf("repository.FindByID id=%s: %w", id, domain.ErrProductNotFound)
	}

	return p, nil
}

func (r *InMemoryProductRepository) List() ([]domain.Product, error) {
	if r.simulateDBError {
		return nil, fmt.Errorf("repository.List: %w", ErrDBConnection)
	}

	r.mu.RLock()
	defer r.mu.RUnlock()

	products := make([]domain.Product, 0, len(r.products))
	for _, p := range r.products {
		products = append(products, p)
	}
	return products, nil
}

func (r *InMemoryProductRepository) Update(product domain.Product) error {
	if r.simulateDBError {
		return fmt.Errorf("repository.Update id=%s: %w", product.ID, ErrDBConnection)
	}

	r.mu.Lock()
	defer r.mu.Unlock()

	if _, exists := r.products[product.ID]; !exists {
		return fmt.Errorf("repository.Update id=%s: %w", product.ID, domain.ErrProductNotFound)
	}

	r.products[product.ID] = product
	return nil
}

func (r *InMemoryProductRepository) Delete(id string) error {
	if r.simulateDBError {
		return fmt.Errorf("repository.Delete id=%s: %w", id, ErrDBConnection)
	}

	r.mu.Lock()
	defer r.mu.Unlock()

	if _, exists := r.products[id]; !exists {
		return fmt.Errorf("repository.Delete id=%s: %w", id, domain.ErrProductNotFound)
	}

	delete(r.products, id)
	return nil
}
