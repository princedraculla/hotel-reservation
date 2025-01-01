package api

import (
	"github.com/gofiber/fiber/v2"
	"github.com/princedraculla/hotel-reservation/db"
	"github.com/princedraculla/hotel-reservation/types"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type BookingHandler struct {
	bookingStore db.BookingStore
}

func NewBookingHandler(bookingStore db.BookingStore) *BookingHandler {
	return &BookingHandler{
		bookingStore: bookingStore,
	}
}

func (bh *BookingHandler) AddBooking(ctx *fiber.Ctx) error {
	var params types.Booking
	if err := ctx.BodyParser(&params); err != nil {
		return err
	}
	roomID, err := primitive.ObjectIDFromHex(ctx.Params("id"))
	if err != nil {
		return err
	}
	user, ok := ctx.Context().Value("user").(*types.User)
	if !ok {
		return ctx.Status(404).JSON("user token not found")
	}
	booking := types.Booking{
		UserID:    user.ID,
		RoomID:    roomID,
		FromDate:  params.FromDate,
		TilDate:   params.TilDate,
		NumPerson: params.NumPerson,
	}
	inserted, err := bh.bookingStore.BookingRoom(ctx.Context(), &booking)
	if err != nil {
		return err
	}
	return ctx.JSON(inserted)
}
