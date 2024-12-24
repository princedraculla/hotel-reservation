package api

import (
	"github.com/gofiber/fiber/v2"
	"github.com/princedraculla/hotel-reservation/db"
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
