package db

func userCollection() {
	UserCollection := db().Database("hotel-reservation").Collection("users")
}
