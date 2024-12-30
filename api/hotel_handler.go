package api

import (
	"github.com/gofiber/fiber/v2"
	"github.com/princedraculla/hotel-reservation/db"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type HotelHandler struct {
	HotelStore db.HotelStore
	RoomeStore db.RoomStore
}

func NewHotelHandler(hotelStore db.HotelStore, roomStore db.RoomStore) *HotelHandler {
	return &HotelHandler{
		HotelStore: hotelStore,
		RoomeStore: roomStore,
	}
}

func (h *HotelHandler) GetHotels(c *fiber.Ctx) error {
	result, err := h.HotelStore.GetHotels(c.Context())
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
	rooms, err := h.RoomeStore.GetRooms(ctx.Context(), filter)
	if err != nil {
		return nil
	}
	return ctx.JSON(rooms)
}
