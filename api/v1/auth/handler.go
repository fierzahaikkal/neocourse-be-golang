package auth

import (
	"encoding/json"
	"net/http"

	"github.com/fierzahaikkal/neocourse-be-golang/internal/model/user"
	"github.com/fierzahaikkal/neocourse-be-golang/internal/usecase"
	"github.com/fierzahaikkal/neocourse-be-golang/pkg/utils"
	log "github.com/sirupsen/logrus"
)

type AuthHandler struct {
	AuthUseCase *usecase.AuthUseCase
	JWTSecret   string
	log         *log.Logger
}

func NewAuthHandler(authUseCase *usecase.AuthUseCase, jwtSecret string) *AuthHandler {
	return &AuthHandler{
		AuthUseCase: authUseCase,
		JWTSecret:   jwtSecret,
		log:         log.New(),
	}
}

func (h *AuthHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Methods", "OPTIONS, GET, POST, PUT")
	w.Header().Set("Access-Control-Allow-Headers", "Content-Type, X-CSRF-Token")

	switch r.URL.Path {
	case "/api/v1/auth/signup":
		if r.Method == http.MethodPost {
			h.SignUp(w, r) // Handle the POST request for SignUp
		} else {
			http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		}
	case "/api/v1/auth/signin":
		if r.Method == http.MethodPost {
			h.SignIn(w, r) // Handle the POST request for SignIn
		} else {
			http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		}
	default:
		http.NotFound(w, r)
	}
}

func (h *AuthHandler) SignUp(w http.ResponseWriter, r *http.Request) {
	var req user.SignUpRequest
	err := json.NewDecoder(r.Body).Decode(&req)
	if err != nil {
		utils.ErrorResponse(w, "Invalid request body", http.StatusBadRequest)
		return
	}

	if err := utils.Validate.Struct(&req); err != nil {
		utils.ErrorResponse(w, err.Error(), http.StatusBadRequest)
		return
	}

	token, err := h.AuthUseCase.SignUp(req, h.JWTSecret)

	statusCode, err := utils.HandleError(err)
	if err != nil {
		utils.ErrorResponse(w, err.Error(), statusCode)
		return
	}
	utils.SuccessResponse(w, map[string]interface{}{"message": "User created successfully", "token": token}, statusCode)
}

func (h *AuthHandler) SignIn(w http.ResponseWriter, r *http.Request) {
	var req user.SignInRequest
	err := json.NewDecoder(r.Body).Decode(&req)
	if err != nil {
		utils.ErrorResponse(w, "Invalid request body", http.StatusBadRequest)
		return
	}

	if err := utils.Validate.Struct(&req); err != nil {
		utils.ErrorResponse(w, err.Error(), http.StatusBadRequest)
		return
	}

	token, err := h.AuthUseCase.SignIn(req, h.JWTSecret)
	statusCode, err := utils.HandleError(err)
	if err != nil {
		utils.ErrorResponse(w, err.Error(), statusCode)
		return
	}
	utils.SuccessResponse(w, map[string]interface{}{"message": "User logged in successfully", "token": token}, http.StatusOK)
}
