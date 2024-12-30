package api

import (
	"github.com/gofiber/fiber/v2"
	"github.com/princedraculla/hotel-reservation/db"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type HotelHandler struct {
	store *db.Store
}

func NewHotelHandler(store *db.Store) *HotelHandler {
	return &HotelHandler{
		store: store,
	}
}

func (h *HotelHandler) GetHotels(c *fiber.Ctx) error {
	result, err := h.store.Hotel.GetHotels(c.Context())
	if err != nil {
		return err
	}
	return c.JSON(result)
}

func (h *HotelHandler) GetRooms(ctx *fiber.Ctx) error {
	id := ctx.Params("id")
	oid, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return err
	}
	filter := bson.M{"hotelID": oid}
	rooms, err := h.store.Rooms.GetRooms(ctx.Context(), filter)
	if err != nil {
		return nil
	}
	return ctx.JSON(rooms)
}
