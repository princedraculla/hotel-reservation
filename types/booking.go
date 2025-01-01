package types

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Booking struct {
	ID        primitive.ObjectID `bson:"_id,omitempty" json:"id,omitempty"`
	UserID    primitive.ObjectID `bson:"userID" json:"userID"`
	RoomID    primitive.ObjectID `bson:"roomID" json:"roomID"`
	NumPerson int                `bson:"numPerson" json:"numPerson"`
	FromDate  time.Time          `bson:"fromDate" json:"fromDate"`
	TilDate   time.Time          `bson:"tilDate" json:"tilDate"`
}

type CreateBookingParams struct {
	NumPerson int       `json:"numPerson"`
	FromDate  time.Time `json:"fromDate"`
	TilDate   time.Time `json:"tilDate"`
}
