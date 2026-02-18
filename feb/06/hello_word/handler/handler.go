package handler

import (
	"context"
	"net/http"
	"fmt"
	"encoding/json"
	kivik "github.com/go-kivik/kivik/v4"
)

type Handler struct {
	couchdb *kivik.DB
}

type heathResponse struct {
	Status  string `json:"status"`
	Message string `json:"message"` 
}

type NotebookCreate struct {
	ID          string `json:"id"`
	Name        string `json:"name"`
	Description string `json:"description"`
}

type Notebook struct {
	ID          string `json:"_id"`
	Rev	    string `json:"_rev"`
	Name        string `json:"name"`
	Description string `json:"description"`
}


//Constructor
func New(c *kivik.DB) *Handler {
	return &Handler{
		couchdb: c,
	}
}

func (h *Handler) Heath(w http.ResponseWriter, r *http.Request) {
	resp := heathResponse{Status: "ok", Message:"Everthing is cool here..."}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)

	if err := json.NewEncoder(w).Encode(resp); err != nil {
		http.Error(w, "Falha ao codificar resposta", http.StatusInternalServerError)
		return
	}
}

func (h *Handler) Create(w http.ResponseWriter, r *http.Request) {
	//Input data
	defer r.Body.Close()
	ctx := context.TODO()

	var body NotebookCreate
	if err := json.NewDecoder(r.Body).Decode(&body); err != nil {
		http.Error(w, "Falha ao criar notebook", http.StatusInternalServerError)
		return
	}

	//Save data in DB
	resp, err := h.couchdb.Put(ctx, body.ID, body)
	if err != nil {
		http.Error(w, "Erro ao savar dados no DB"+err.Error(), http.StatusInternalServerError)
		return
	}

	fmt.Println("couch resp: ", resp)

	//Return infos
	if err := json.NewEncoder(w).Encode(resp); err != nil {
		http.Error(w, "Erro ao retornar dados salvos", http.StatusInternalServerError)
		return
	}
}

func (h *Handler) Get(w http.ResponseWriter, r *http.Request) {
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

	rows, err := h.couchdb.AllDocs(ctx, kivik.Options{})
	if err != nil {
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
