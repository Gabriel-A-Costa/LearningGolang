package service

import (
	"context"
	"github.com/google/uuid"
	"github.com/gabriel-a-costa/LearningGolang/feb/06/first_service/repository"
	"github.com/gabriel-a-costa/LearningGolang/feb/06/first_service/model"
	kivik "github.com/go-kivik/kivik/v4"
)

type Service struct {
	repo *repository.Repository	
}

func NewService(r *repository.Repository) *Service {
	return &Service {
		repo: r,
	}
}

func (s *Service) Create(ctx context.Context, input model.CreateNotebookInputDTO) (*Notebook, error) {

	id := uuid.NewString()
	err := validateNotebookName(input.Name)
	if err != nil {
		return nil, err
	}

	s.repo.Create(ctx, new_notebook)

	new_notebook := &Notebook{
		ID: id,
		Name: input.Name,
		Description: input.Description,
	}	

	return new_notebook, nil
}

func (s *Service) Get(w http.ResponseWriter, r *http.Request) {
	defer r.Body.Close()
	ctx := context.TODO()

	var nb Notebook
	id := r.PathValue("id")

	err := h.couchdb.Get(ctx, id).ScanDoc(&nb)
	if err != nil {
		http.Error(w, "Erro ao buscar notebook"+err.Error(), http.StatusInternalServerError)
		return
	}

	if err := json.NewEncoder(w).Encode(nb); err != nil {
		http.Error(w, "Erro ao retornar notebook", http.StatusInternalServerError)
		return
	}
}

func (h *Handler) List(w http.ResponseWriter, r *http.Request) {
	defer r.Body.Close()
	ctx := context.TODO()

	rows := h.couchdb.AllDocs(ctx)
	if err := rows.Err(); err != nil {
		http.Error(w, "Erro ao retornar lista de notebooks"+err.Error(), http.StatusInternalServerError)
		return
	}

	var nbs []Notebook
	for rows.Next() {
		var nb Notebook
		if err := rows.ScanDoc(&nb); err != nil {
			http.Error(w, "Erro ao Decodificar notebook"+err.Error(), http.StatusInternalServerError)
		}
		nbs = append(nbs, nb)
	}

	if err := json.NewEncoder(w).Encode(nbs); err != nil {
		http.Error(w, "Erro ao exibir lista"+err.Error(), http.StatusInternalServerError)
		return
	}
}


func (h *Handler) Update(w http.ResponseWriter, r *http.Request) {
	defer r.Body.Close()
	ctx := context.TODO()

	var body Notebook
	id := r.PathValue("id")

	if err := json.NewDecoder(r.Body).Decode(&body); err != nil {
		http.Error(w, "Erro na leitura do Body!", http.StatusInternalServerError)
		return	
	}

	resp, err := h.couchdb.Put(ctx, id, body)
	if err != nil {
		http.Error(w, "Erro ao atualizar o notebook"+err.Error(), http.StatusInternalServerError)
		return
	}

	fmt.Println("couch response: ", resp)

	if err := json.NewEncoder(w).Encode(resp); err != nil {
		http.Error(w, "Erro ao exibir dados", http.StatusInternalServerError)
		return
	}
}

func (h *Handler) Delete(w http.ResponseWriter, r *http.Request) {
	defer r.Body.Close()
	ctx := context.TODO()

	id := r.PathValue("id")
	rev := r.PathValue("rev")

	_, err := h.couchdb.Delete(ctx, id, rev)
	if err != nil {
		http.Error(w, "Erro ao apagar notebook"+err.Error(), http.StatusInternalServerError)
		return
	}

	if err := json.NewEncoder(w).Encode("Notebook Apagado Com Sucesso!"); err != nil {
		http.Error(w, "Erro ao exibir mensagem", http.StatusInternalServerError)
		return
	}
}
