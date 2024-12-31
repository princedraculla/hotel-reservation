package middleware

import (
	"os"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

func JWTAuthentication() {

}

func CreateToken(email string, password string) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodES256, jwt.MapClaims{
		"email":    email,
		"password": password,
		"exp":      time.Now().Add(time.Hour * 24).Unix(),
	})
	tokenStr, err := token.SignedString(os.Getenv("secret_key"))
	if err != nil {
		return "", nil
	}
	return tokenStr, nil
}
