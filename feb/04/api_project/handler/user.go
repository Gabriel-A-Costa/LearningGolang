package handler

import (
	"encoding/json"
	"net/http"
	"strings"

	"github.com/gabriel-a-costa/LearningGolang/feb/04/api_project/service"
)

type UserResponse struct {
	ID   string `json:"id"`
	Name string `json:"name"`
	Age  int    `json:"age"`
}

func ListUsers(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	users := service.GetAllUsers()

	var result []UserResponse

	for _, user := range users {
		result = append(result, UserResponse{
			ID:   user.ID,
			Name: user.Name,
			Age:  user.Age,
		})
	}

	if err := json.NewEncoder(w).Encode(result); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}

func GetUserById(w http.ResponseWriter, r *http.Request) {
	id := strings.TrimPrefix(r.URL.Path, "/users/")
	user, err := service.GetUserById(id)
	if err != nil {
		http.Error(w, err.Error(), http.StatusNotFound)
		return
	}

	result := UserResponse{
		ID:   user.ID,
		Name: user.Name,
		Age:  user.Age,
	}

	w.Header().Set("Content-Type", "application/json")

	if err := json.NewEncoder(w).Encode(result); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}
