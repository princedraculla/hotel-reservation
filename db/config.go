package db

const DBNAME = "hotel-reservation"
const TestDBNAME = "hotel-reservation-test"
const DBURI = "mongodb://localhost:27017"

type Store struct {
	User  UserStore
	Hotel HotelStore
	Rooms RoomStore
}
