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

func (uc *AuthUseCase) SignUp(req *user.SignUpRequest) (*entity.User, error) {
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(req.Password), bcrypt.DefaultCost)
	if err != nil {
		return nil, err
	}

	user := entity.User{
		ID:       utils.GenUUID(),
		Username: req.Username,
		Email:    req.Email,
		Name:     req.Name,
		Password: string(hashedPassword),
	}

	err = uc.UserRepo.Register(&user)
	if err != nil {
		return nil, err
	}

	return &user, nil
}

func (uc *AuthUseCase) SignIn(req *user.SignInRequest) (*entity.User, error) {
	userFromDb, err := uc.UserRepo.FindByEmail(req.Email)
	if err != nil {
		return nil, err
	}

	err = bcrypt.CompareHashAndPassword([]byte(userFromDb.Password), []byte(req.Password))
	if err != nil {
		return nil, err
	}

	return userFromDb, nil
}

func (uc *AuthUseCase) GetUser(id string) (*entity.User, error) {
	userFromDB, err := uc.UserRepo.FindByID(id)
	if err != nil {
		return nil, err
	}

	return userFromDB, nil
}