package api

import (
	"github.com/gofiber/fiber/v2"
	"github.com/princedraculla/hotel-reservation/types"
)

func HandleGetUser(c *fiber.Ctx) error {
	user := types.User{
		FirstName: "amir",
		LastName:  "torkashvand",
	}

	return c.JSON(user)
}
