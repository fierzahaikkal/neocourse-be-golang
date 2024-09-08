package utils

import (
	"time"

	"github.com/fierzahaikkal/neocourse-be-golang/internal/entity"
	"github.com/golang-jwt/jwt"
)

func GenerateJWT(user *entity.User, secret string) (string, error) {
	claims := jwt.MapClaims{
		"email": user.Email,
		"exp":   time.Now().Add(time.Hour * 72).Unix(), // Token expiration time
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString([]byte(secret))
}
