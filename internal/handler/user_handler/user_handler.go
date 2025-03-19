package user_handler

import (
	"adp_practice1/internal/domain"
	"adp_practice1/internal/usecase"
	"encoding/json"
	"github.com/gorilla/mux"
	"net/http"
	"strconv"
)

type UserHandler struct {
	userUseCase *usecase.UserUsecase
}

func NewUserHandler(userUseCase *usecase.UserUsecase) *UserHandler {
	return &UserHandler{userUseCase: userUseCase}
}

func (h *UserHandler) CreateUserHandler(w http.ResponseWriter, r *http.Request) {
	var user domain.User
	err := json.NewDecoder(r.Body).Decode(&user)
	if err != nil {
		http.Error(w, "Invalid input", http.StatusBadRequest)
		return
	}

	createdUser, err := h.userUseCase.CreateUser(user.Name, user.Email)
	if err != nil {
		http.Error(w, "Failed to create user", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(createdUser)
}

func (h *UserHandler) GetUserHandler(w http.ResponseWriter, r *http.Request) {
	id, err := parseIDParam(r)
	if err != nil {
		http.Error(w, "Invalid user ID", http.StatusBadRequest)
		return
	}

	user, err := h.userUseCase.GetUserByID(id)
	if err != nil {
		http.Error(w, "User not found", http.StatusNotFound)
		return
	}

	json.NewEncoder(w).Encode(user)
}

func parseIDParam(r *http.Request) (uint64, error) {
	vars := mux.Vars(r)
	return strconv.ParseUint(vars["id"], 10, 64)
}
