package auth

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/fierzahaikkal/neocourse-be-golang/internal/repository"
	"github.com/fierzahaikkal/neocourse-be-golang/internal/entity"
	"github.com/fierzahaikkal/neocourse-be-golang/pkg/utils"
	"golang.org/x/crypto/bcrypt"
)

type AuthHandler struct {
	UserRepo  *repository.UserRepository
	JWTSecret string
}

func NewAuthHandler(userRepo *repository.UserRepository, jwtSecret string) *AuthHandler {
	return &AuthHandler{
		UserRepo:  userRepo,
		JWTSecret: jwtSecret,
	}
}

func (h *AuthHandler) SignUp(w http.ResponseWriter, r *http.Request) {
	var req utils.SignUpRequest
	err := json.NewDecoder(r.Body).Decode(&req)
	if err != nil {
		utils.ErrorResponse(w, "Invalid request body", http.StatusBadRequest)
		return
	}

	if err := req.Validate(); err != nil {
		utils.ErrorResponse(w, err.Error(), http.StatusBadRequest)
		return
	}

	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(req.Password), bcrypt.DefaultCost)
	if err != nil {
		log.Printf("Error hashing password: %v", err)
		utils.ErrorResponse(w, "Server error", http.StatusInternalServerError)
		return
	}

	user := entity.User{
		Email:    req.Email,
		Password: string(hashedPassword),
	}

	err = h.UserRepo.CreateUser(&user)
	if err != nil {
		utils.ErrorResponse(w, "User already exists", http.StatusConflict)
		return
	}

	utils.SuccessResponse(w, "User created successfully")
}

func (h *AuthHandler) SignIn(w http.ResponseWriter, r *http.Request) {
	var req utils.SignUpRequest
	err := json.NewDecoder(r.Body).Decode(&req)
	if err != nil {
		utils.ErrorResponse(w, "Invalid request body", http.StatusBadRequest)
		return
	}

	if err := req.Validate(); err != nil {
		utils.ErrorResponse(w, err.Error(), http.StatusBadRequest)
		return
	}

	user, err := h.UserRepo.FindByEmail(req.Email)
	if err != nil {
		utils.ErrorResponse(w, "Invalid email or password", http.StatusUnauthorized)
		return
	}

	err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(req.Password))
	if err != nil {
		utils.ErrorResponse(w, "Invalid email or password", http.StatusUnauthorized)
		return
	}

	token, err := utils.GenerateJWT(user, h.JWTSecret)
	if err != nil {
		utils.ErrorResponse(w, "Failed to generate token", http.StatusInternalServerError)
		return
	}

	utils.SuccessResponse(w, map[string]string{"token": token})
}
