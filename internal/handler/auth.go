package handler

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/Oliver1ck/testApi/internal/models"
	"github.com/Oliver1ck/testApi/internal/repository"
	"golang.org/x/crypto/bcrypt"
)

type AuthHandler struct {
	userRepo repository.UserRepository
}

// Конструктор
func NewAuthHandler(userRepo repository.UserRepository) *AuthHandler {
	return &AuthHandler{
		userRepo: userRepo,
	}
}

type responseUser struct {
	UserName string `json:"username"`
	Email    string `json:"email"`
	Password string `json:"password"`
}

func (h *AuthHandler) RegisterHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		w.Header().Set("Allow", "POST, OPTIONS")
		http.Error(w, "Method Not Allowed", http.StatusMethodNotAllowed)
		return
	}

	var req responseUser
	err := json. NewDecoder(r.Body).Decode(&req)
	if err != nil {
		http.Error(w, fmt.Sprintf("ошибка декодирования:  %v", err), http.StatusBadRequest)
		return 
	}

	if req.UserName == "" || req.Email == "" || req.Password == "" {
		http.Error(w, "все поля обязательны", http.StatusBadRequest)
		return
	}

	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(req.Password), bcrypt.DefaultCost)
	if err != nil {
		http.Error(w, "ошибка хеширования пароля", http.StatusInternalServerError)
		return
	}

	user := &models. User{
		Username: req. UserName,
		Email:    req.Email,
		Password: string(hashedPassword),
	}

	createdUser, err := h.userRepo.CreateUser(user)
	if err != nil {
		http.Error(w, fmt.Sprintf("ошибка создания пользователя: %v", err), http.StatusInternalServerError)
		return
	}

	response := map[string]interface{}{
		"id":         createdUser.ID,
		"username":   createdUser.Username,
		"email":      createdUser.Email,
		"created_at":  createdUser.CreatedAt,
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(response)
}