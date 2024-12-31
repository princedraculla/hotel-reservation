package api

import (
	"fmt"
	"net/http"
	"os"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v5"
	"github.com/princedraculla/hotel-reservation/db"
	"github.com/princedraculla/hotel-reservation/types"
)

var (
	secretKey = os.Getenv("secret_key")
)

type AuthHandler struct {
	userStore db.UserStore
}

func NewAuthHandler(userStore db.UserStore) *AuthHandler {
	return &AuthHandler{
		userStore: userStore,
	}
}

type AuthParams struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

type AuthResponse struct {
	User  *types.User `json:"user"`
	Token string      `json:"token"`
}

type genericResp struct {
	Type string `json:"type"`
	Msg  string `json:"msg"`
}

func InvalidCredentials(ctx *fiber.Ctx) error {
	return ctx.Status(http.StatusBadRequest).JSON(genericResp{
		Type: "error",
		Msg:  "invalid credentials",
	})
}

func (ah *AuthHandler) HandleAutheticate(ctx *fiber.Ctx) error {
	var params AuthParams
	if err := ctx.BodyParser(&params); err != nil {
		return err
	}
	user, err := ah.userStore.GetUserByEmail(ctx.Context(), params.Email)
	if err != nil {
		return err
	}
	if !types.IsValidPassword(user.EncryptedPassword, params.Password) {
		return InvalidCredentials(ctx)
	}
	resp := AuthResponse{
		User:  user,
		Token: CreateToken(user),
	}
	return ctx.JSON(resp)
}

func CreateToken(user *types.User) string {

	claims := jwt.MapClaims{
		"id":    user.ID,
		"email": user.Email,
		"exp":   time.Now().Add(time.Hour * 2).Unix(),
	}
	toekn := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenStr, err := toekn.SignedString([]byte(secretKey))
	if err != nil {
		fmt.Println("faild to sign token with secret", err)
	}
	return tokenStr
}
