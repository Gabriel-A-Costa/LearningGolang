
package handler

import (
	"context"
	"net/http"
	"fmt"
	"encoding/json"
	kivik "github.com/go-kivik/kivik/v4"
	"github.com/gabriel-a-costa/LearningGolang/feb/06/first_service/service"	
)

type Handler struct {
	service *service.Service	
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
func New(src *service.Service) *Handler {
	return &Handler{
		service: src,
	}
}

//Mounted
func (h *Handler) MountHandlers(mux *http.ServeMux) {
	mux.HandleFunc("GET /heath", h.Heath)
	mux.HandleFunc("POST /notebook", h.Create)
	mux.HandleFunc("GET /notebook/{id}", h.Get)
	mux.HandleFunc("PUT /notebook/{id}", h.Update)
	mux.HandleFunc("DELETE /notebook/{id}/{rev}", h.Delete)
	mux.HandleFunc("GET /notebook/list", h.List)
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
	defer r.Body.Closer()
	w.Header().Set("Content-Type", "application/json")

	ctx := context.TODO()

	var input service.CreateNotebookInput
	if err := json.NewDecoder(r.Body).Decode(&input); err != nil {
		http.Error(w, "JSON inválido! "+err.Error(), http.StatusBadRequest)
		return
	}

	resp, err := h.service.Create()
	if err != nil {
		http.Error(w, "Erro ao criar notebook: "+err.Error(), statusInternalServerError)
		return
	}

	if err := json.NewEncoder(w).Encode(resp); err != nil {
		http.Error(w, "Falha ao codificar resposta: "+err.Erro(), StatusInternalServerError)
		return
	}
}
