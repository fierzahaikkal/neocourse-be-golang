package utils

import (
	"fmt"
	"time"

	"github.com/fierzahaikkal/neocourse-be-golang/internal/entity"
	"github.com/golang-jwt/jwt"
)

// GenerateJWT generates a JWT token for a given user
func GenerateJWT(user *entity.User, secret string) (string, error) {
	claims := jwt.MapClaims{
		"id": user.ID,
		"email": user.Email,
		"exp":   time.Now().Add(time.Hour * 4).Unix(),
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString([]byte(secret))
}

// ValidateJWT validates a JWT token
func ValidateJWT(tokenString string, secret string) (bool, error) {
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}
		return []byte(secret), nil
	})
	if err != nil {
		return false, err
	}
	return token.Valid, nil
}

// ExtractEmailFromJWT extracts the email from a JWT token
func ExtractEmailFromJWT(tokenString string) (string, error) {
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		return []byte("secret"), nil // Use your secret key here
	})
	if err != nil {
		return "", err
	}
	claims, ok := token.Claims.(jwt.MapClaims)
	if !ok || !token.Valid {
		return "", fmt.Errorf("invalid token")
	}
	email, ok := claims["email"].(string)
	if !ok {
		return "", fmt.Errorf("email claim is not a string")
	}
	return email, nil
}

func ExtractIDFromJWT(tokenString string) (string, error) {
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		return []byte("secret"), nil // Use your secret key here
	})
	if err != nil {
		return "", err
	}
	claims, ok := token.Claims.(jwt.MapClaims)
	if !ok || !token.Valid {
		return "", fmt.Errorf("invalid token")
	}
	id, ok := claims["id"].(string)
	if !ok {
		return "", fmt.Errorf("email claim is not a string")
	}
	return id, nil
}
