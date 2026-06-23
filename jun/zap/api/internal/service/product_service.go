package service

import (
	"fmt"

	"github.com/google/uuid"

	"github.com/gabriel-a-costa/LearningGolang/jun/zap/api/internal/domain"
)

// ProductService orquestra a lógica de negócio.
// Recebe domain.Repository (interface), não o tipo concreto.
// Isso permite testar o service com um mock sem banco de dados.
type ProductService struct {
	repo domain.Repository
}

func New(repo domain.Repository) *ProductService {
	return &ProductService{repo: repo}
}

// ══════════════════════════════════════════════════════════════════════════════
// REGRAS DA CAMADA SERVICE
//
// RESPONSABILIDADE: regras de negócio, validação, orquestração.
//
// O QUE RETORNAR:
//   - Erros de validação: fmt.Errorf("service.Op: %w", &domain.ValidationError{...})
//   - Erros do repository: fmt.Errorf("service.Op: %w", err)  ← adiciona contexto
//
// O QUE NÃO LOGAR:
//   - NADA. O service propaga, não decide.
//
// POR QUE NÃO LOGAR?
//   Regra de ouro: "Log onde você TRATA o erro, não onde você o PROPAGA."
//   O service propaga erros para o handler decidir o que fazer.
//   Se logar aqui E o handler logar → log duplicado para o mesmo erro.
//
// ERRO COMUM:
//   fmt.Errorf("service: %v", err)  ← %v QUEBRA errors.Is() upstream!
//   fmt.Errorf("service: %w", err)  ← %w PRESERVA a cadeia. Sempre use %w.
//
// TRILHA DE AUDITORIA via context no erro:
//   Depois de 3 camadas, o erro parece:
//   "service.GetByID: repository.FindByID id=abc: product not found"
//   Você vê exatamente onde nasceu o erro sem precisar de stack trace.
// ══════════════════════════════════════════════════════════════════════════════

func (s *ProductService) Create(input domain.CreateInput) (domain.Product, error) {
	// Cenário 3: validar ANTES de ir ao banco.
	// Por que no service e não no handler?
	// Validação é regra de negócio. Se você mover Create() para um worker
	// ou CLI, a validação continua funcionando automaticamente.
	if err := input.Validate(); err != nil {
		// input.Validate() retorna *domain.ValidationError
		// Aqui adicionamos contexto do service mas NÃO logamos.
		// O handler receberá e verificará: errors.Is(err, domain.ErrInvalidProduct) → true
		return domain.Product{}, fmt.Errorf("service.Create: %w", err)
	}

	product := domain.Product{
		ID:    uuid.New().String(),
		Name:  input.Name,
		Price: input.Price,
		Stock: input.Stock,
	}

	if err := s.repo.Create(product); err != nil {
		// Não sabemos se é ErrProductDuplicate ou ErrDBConnection — não precisamos saber.
		// Adicionamos contexto e propagamos. O handler decide o HTTP status.
		return domain.Product{}, fmt.Errorf("service.Create: %w", err)
	}

	return product, nil
}

func (s *ProductService) GetByID(id string) (domain.Product, error) {
	// Cenário 3: validação simples antes do banco
	if id == "" {
		return domain.Product{}, fmt.Errorf("service.GetByID: %w",
			&domain.ValidationError{Field: "id", Message: "é obrigatório"},
		)
	}

	product, err := s.repo.FindByID(id)
	if err != nil {
		// Cadeia de erro neste ponto:
		//   "repository.FindByID id=abc: product not found"
		// Depois do wrap:
		//   "service.GetByID: repository.FindByID id=abc: product not found"
		// errors.Is(err, domain.ErrProductNotFound) → ainda true ✓
		return domain.Product{}, fmt.Errorf("service.GetByID: %w", err)
	}

	return product, nil
}

func (s *ProductService) List() ([]domain.Product, error) {
	products, err := s.repo.List()
	if err != nil {
		return nil, fmt.Errorf("service.List: %w", err)
	}
	return products, nil
}

func (s *ProductService) Update(id string, input domain.UpdateInput) (domain.Product, error) {
	if id == "" {
		return domain.Product{}, fmt.Errorf("service.Update: %w",
			&domain.ValidationError{Field: "id", Message: "é obrigatório"},
		)
	}
	if err := input.Validate(); err != nil {
		return domain.Product{}, fmt.Errorf("service.Update: %w", err)
	}

	product := domain.Product{
		ID:    id,
		Name:  input.Name,
		Price: input.Price,
		Stock: input.Stock,
	}

	if err := s.repo.Update(product); err != nil {
		return domain.Product{}, fmt.Errorf("service.Update: %w", err)
	}

	return product, nil
}

func (s *ProductService) Delete(id string) error {
	if id == "" {
		return fmt.Errorf("service.Delete: %w",
			&domain.ValidationError{Field: "id", Message: "é obrigatório"},
		)
	}

	if err := s.repo.Delete(id); err != nil {
		return fmt.Errorf("service.Delete: %w", err)
	}

	return nil
}
