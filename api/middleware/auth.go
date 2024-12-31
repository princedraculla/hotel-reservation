package middleware

import (
	"fmt"
	"os"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

var (
	secretKey = os.Getenv("secret_key")
)

func JWTAuthentication() {

}

func CreateToken(email string, password string) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodES256, jwt.MapClaims{
		"email":    email,
		"password": password,
		"exp":      time.Now().Add(time.Hour * 24).Unix(),
	})
	tokenStr, err := token.SignedString(secretKey)
	if err != nil {
		return "", nil
	}
	return tokenStr, nil
}

func VerifyToken(tokenStr string) error {
	token, err := jwt.Parse(tokenStr, func(t *jwt.Token) (interface{}, error) {
		return secretKey, nil
	})
	if err != nil {
		return err
	}
	if !token.Valid {
		return fmt.Errorf("toekn its not value %s", token)
	}
	return nil
}
