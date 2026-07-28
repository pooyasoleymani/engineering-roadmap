package utils

import (
	"time"

	"github.com/golang-jwt/jwt/v5"
)

func GenerateToken(id int64, secretKey string) (string, error) {
	token := jwt.NewWithClaims(
		jwt.SigningMethodHS256, jwt.MapClaims{
			"user_id": id,
			"exp":     time.Now().Add(5 * time.Minute).Unix(),
		},
	)
	tokenString, err := token.SignedString([]byte(secretKey))
	if err != nil {
		return "", err
	}
	return tokenString, nil
}
