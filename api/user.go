package api

import (
	"context"

	"github.com/gofiber/fiber/v2"
	"github.com/princedraculla/hotel-reservation/db"
	"github.com/princedraculla/hotel-reservation/types"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type UserHandler struct {
	userStore db.UserStore
}

func NewUserHandler(userStore db.UserStore) *UserHandler {
	return &UserHandler{
		userStore: userStore,
	}
}

func (h *UserHandler) HandlePostUser(c *fiber.Ctx) error {
	var params types.CreateUserParams
	if err := c.BodyParser(&params); err != nil {
		return err
	}
	if errors := types.CreateUserParams.InputValidation(params); len(errors) > 0 {
		return c.JSON(errors)
	}
	hashedUser, err := types.EncodingUserPassword(params)
	if err != nil {
		return err
	}
	insertedUser, err := h.userStore.AddUser(c.Context(), hashedUser)
	if err != nil {
		return err
	}
	return c.JSON(insertedUser)
}

func (h *UserHandler) HandleGetUser(c *fiber.Ctx) error {
	id := c.Params("id")
	ctx := context.Background()
	user, err := h.userStore.GetUserByID(ctx, id)
	if err != nil {
		return err
	}
	return c.JSON(user)
}
func (h *UserHandler) HandleGetUsers(c *fiber.Ctx) error {

	userList, err := h.userStore.UserList(c.Context())

	if err != nil {
		return err
	}

	return c.JSON(userList)
}

func (h *UserHandler) HandleDeleteUser(c *fiber.Ctx) error {
	userID := c.Params("id")
	err := h.userStore.DeleteUser(c.Context(), userID)
	if err != nil {
		return err
	}
	return nil
}

func (h *UserHandler) HandleUserUpdate(c *fiber.Ctx) error {
	var (
		params types.UpdateUserParams
		userID = c.Params("id")
	)
	if err := c.BodyParser(&params); err != nil {
		return err
	}

	oid, err := primitive.ObjectIDFromHex(userID)
	if err != nil {
		return err
	}
	filter := bson.D{{Key: "_id", Value: oid}}

	if err := h.userStore.UpdateUser(c.Context(), filter, params); err != nil {
		return err
	}

	return c.JSON(userID)
}
