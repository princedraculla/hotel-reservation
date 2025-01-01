package types

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Booking struct {
	UserID    primitive.ObjectID `bson:"userID" json:"userID"`
	RoomID    primitive.ObjectID `bson:"roomID" json:"roomID"`
	NumPerson int                `bson:"numPerson" json:"numPerson"`
	FromDate  time.Time          `bson:"fromDate" json:"fromDate"`
	TilDate   time.Time          `bson:"tilDate" json:"tilDate"`
}
