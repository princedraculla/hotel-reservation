package api

import (
	"fmt"
	"net/http"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v5"
	"github.com/princedraculla/hotel-reservation/db"
)

func JWTAuthenticate(userStore db.UserStore) fiber.Handler {
	return func(ctx *fiber.Ctx) error {
		token := ctx.Get("Authorization")
		if token == "" {
			fmt.Println("token not found")
			return ErrUnAuthorized()
		}
		claims, err := ValidateToken(token)
		if err != nil {
			return err
		}
		// converting time.duration to float type to can conver to int
		expires := int64(claims["exp"].(float64))
		if time.Now().Unix() > expires {
			return NewError(http.StatusUnauthorized, "token expired & sign in first")
		}
		userID := claims["id"].(string)
		user, err := userStore.GetUserByID(ctx.Context(), userID)
		if err != nil {
			return ErrUnAuthorized()
		}
		ctx.Context().SetUserValue("user", user)
		return ctx.Next()

	}
}

func ValidateToken(tokenStr string) (jwt.MapClaims, error) {
	token, err := jwt.Parse(tokenStr, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			fmt.Println("invalid signing method", token.Header["alg"])
			return nil, ErrUnAuthorized()
		}
		return []byte(secretKey), nil
	})
	if err != nil {
		fmt.Println("faild to parse token", err)
		return nil, ErrUnAuthorized()
	}
	if !token.Valid {
		fmt.Println("invalid token")
		return nil, ErrUnAuthorized()
	}
	claims, ok := token.Claims.(jwt.MapClaims)
	if !ok {
		return nil, ErrUnAuthorized()
	}
	return claims, nil
}
