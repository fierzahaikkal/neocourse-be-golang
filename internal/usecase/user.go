package usecase

import (
	"github.com/fierzahaikkal/neocourse-be-golang/internal/entity"
	"github.com/fierzahaikkal/neocourse-be-golang/internal/repository"
	"github.com/fierzahaikkal/neocourse-be-golang/pkg/utils"

	"github.com/fierzahaikkal/neocourse-be-golang/internal/model/user"
	log "github.com/sirupsen/logrus"
	"golang.org/x/crypto/bcrypt"
)

type AuthUseCase struct {
	UserRepo *repository.UserRepository
	log      *log.Logger
}

func NewAuthUseCase(userRepo *repository.UserRepository, log *log.Logger) *AuthUseCase {
	return &AuthUseCase{UserRepo: userRepo, log: log}
}

func (uc *AuthUseCase) SignUp(req user.SignUpRequest, jwtSecret string) (string, error) {
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(req.Password), bcrypt.DefaultCost)
	if err != nil {
		log.Printf("Error hashing password: %v", err)
		return "", err
	}

	user := entity.User{
		ID:       utils.GenUUID(),
		Email:    req.Email,
		Password: string(hashedPassword),
		Username: req.Username,
		Name:     req.Name,
	}

	if err := uc.UserRepo.Register(&user); err != nil {
		return "", err
	}

	// generate token
	token, err := utils.GenerateJWT(&user, jwtSecret)
	if err != nil {
		if err == utils.ErrUserExists {
			return "", utils.ErrUserExists
		}
		return "", err
	}

	return token, nil
}

func (uc *AuthUseCase) SignIn(req user.SignInRequest, jwtSecret string) (string, error) {
	var user entity.User
	userFromDb, err := uc.UserRepo.FindByEmail(req.Email, &user)

	if err != nil {
		if err == utils.ErrRecordNotFound {
			return "", utils.ErrRecordNotFound
		}
		return "", err
	}

	err = bcrypt.CompareHashAndPassword([]byte(userFromDb.Password), []byte(req.Password))
	if err != nil {
		return "", utils.ErrInvalidCredentials
	}

	token, err := utils.GenerateJWT(&user, jwtSecret)
	if err != nil {
		if err == utils.ErrUserExists {
			return "", utils.ErrUserExists
		}
		return "", err
	}

	return token, nil
}
