package domain

import (
	"errors"
	"fmt"
)

// ══════════════════════════════════════════════════════════════════════════════
// ERROS SENTINEL — o "vocabulário" do seu domínio
//
// Por que usar erros sentinel e não strings?
//
//   ✗ Errado:  return errors.New("product not found")  ← cria uma nova instância a cada chamada
//   ✓ Correto: return fmt.Errorf("repo: %w", ErrProductNotFound) ← preserva a identidade
//
// Com sentinel + %w, qualquer camada acima pode fazer:
//   errors.Is(err, ErrProductNotFound) → true
// mesmo que o erro tenha sido embrulhado 3 vezes ao longo do call stack.
//
// Por que definir aqui (domínio) e não no repository?
// O handler precisa checar esses erros para decidir o HTTP status.
// Se os erros fossem definidos no repository, o handler teria que importar
// o repository → quebra a separação de camadas e cria acoplamento errado.
// ══════════════════════════════════════════════════════════════════════════════

var (
	ErrProductNotFound  = errors.New("product not found")
	ErrProductDuplicate = errors.New("product already exists")
	ErrInvalidProduct   = errors.New("invalid product data")
)

// ValidationError é um erro TIPADO que carrega detalhes sobre qual campo falhou.
//
// Quando usar errors.Is() vs errors.As():
//
//   errors.Is(err, ErrInvalidProduct)
//     → "existe ErrInvalidProduct em algum lugar da cadeia?"
//     → use quando você só precisa SABER se foi erro de validação
//
//   var ve *ValidationError
//   errors.As(err, &ve)
//     → "consigo extrair um *ValidationError da cadeia?"
//     → use quando você precisa dos CAMPOS (Field, Message) para resposta detalhada
//
// O Unwrap() abaixo faz os dois funcionarem ao mesmo tempo:
//   errors.Is(err, ErrInvalidProduct) → true   (via Unwrap)
//   errors.As(err, &ve)               → true   (via type assertion)
type ValidationError struct {
	Field   string
	Message string
}

func (e *ValidationError) Error() string {
	return fmt.Sprintf("campo '%s': %s", e.Field, e.Message)
}

// Unwrap encadeia ValidationError ao sentinel ErrInvalidProduct.
// Isso permite que errors.Is(err, ErrInvalidProduct) funcione mesmo quando
// o erro concreto é *ValidationError.
func (e *ValidationError) Unwrap() error {
	return ErrInvalidProduct
}

// ─── Entidade ───────────────────────────────────────────────────────────────

type Product struct {
	ID    string  `json:"id"`
	Name  string  `json:"name"`
	Price float64 `json:"price"`
	Stock int     `json:"stock"`
}

// ─── DTOs de entrada ────────────────────────────────────────────────────────

type CreateInput struct {
	Name  string  `json:"name"`
	Price float64 `json:"price"`
	Stock int     `json:"stock"`
}

// Validate verifica os dados antes de chegar ao repository.
//
// Por que a validação fica no domínio e não no handler?
// Se amanhã você mover esta operação para um worker ou CLI,
// a validação continua funcionando sem precisar duplicar código.
// Validação é regra de NEGÓCIO, não regra de HTTP.
func (c CreateInput) Validate() error {
	if c.Name == "" {
		return &ValidationError{Field: "name", Message: "é obrigatório"}
	}
	if c.Price <= 0 {
		return &ValidationError{Field: "price", Message: fmt.Sprintf("deve ser positivo, recebido %.2f", c.Price)}
	}
	if c.Stock < 0 {
		return &ValidationError{Field: "stock", Message: fmt.Sprintf("não pode ser negativo, recebido %d", c.Stock)}
	}
	return nil
}

type UpdateInput struct {
	Name  string  `json:"name"`
	Price float64 `json:"price"`
	Stock int     `json:"stock"`
}

func (u UpdateInput) Validate() error {
	if u.Name == "" {
		return &ValidationError{Field: "name", Message: "é obrigatório"}
	}
	if u.Price <= 0 {
		return &ValidationError{Field: "price", Message: fmt.Sprintf("deve ser positivo, recebido %.2f", u.Price)}
	}
	if u.Stock < 0 {
		return &ValidationError{Field: "stock", Message: fmt.Sprintf("não pode ser negativo, recebido %d", u.Stock)}
	}
	return nil
}

// ─── Interface do Repository ─────────────────────────────────────────────────
//
// A interface fica no CONSUMIDOR (service), não no PRODUTOR (repository).
// Isso segue o Princípio da Inversão de Dependência e a recomendação da Uber:
// "Accept interfaces, return structs."
//
// Benefício: o service pode ser testado com um mock do Repository
// sem precisar de um banco de dados real.
type Repository interface {
	Create(product Product) error
	FindByID(id string) (Product, error)
	List() ([]Product, error)
	Update(product Product) error
	Delete(id string) error
}
