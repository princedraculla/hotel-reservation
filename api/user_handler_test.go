package api

import (
	"bytes"
	"context"
	"encoding/json"
	"net/http/httptest"
	"testing"

	"github.com/gofiber/fiber/v2"
	"github.com/princedraculla/hotel-reservation/db"
	"github.com/princedraculla/hotel-reservation/types"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

const testdburi = "mongodb://localhost:27017"
const dbname = "hotel-reservation-test"

type testdb struct {
	db.UserStore
}

func (tdb *testdb) teardown(t *testing.T) {
	if err := tdb.UserStore.Drop(context.TODO()); err != nil {
		t.Fatal(err)
	}
}

func setup(t *testing.T) *testdb {
	client, err := mongo.Connect(context.TODO(), options.Client().ApplyURI(testdburi))
	if err != nil {
		t.Fatal(err)
	}
	return &testdb{
		UserStore: db.NewMongoUserStorer(client, dbname),
	}
}

func TestPostUser(t *testing.T) {
	tdb := setup(t)

	defer tdb.teardown(t)

	app := fiber.New()
	userhandler := NewUserHandler(tdb.UserStore)
	app.Post("/", userhandler.HandlePostUser)

	params := types.CreateUserParams{
		FirstName: "amir",
		LastName:  "torkashvand",
		Email:     "amir@gmail.com",
		Password:  "amir1234",
	}
	b, err := json.Marshal(params)
	if err != nil {
		t.Error(err)
	}

	req := httptest.NewRequest("POST", "/", bytes.NewReader(b))
	req.Header.Add("Content-Type", "application/json")

	resp, err := app.Test(req)
	if err != nil {
		t.Error(err)
	}

	var user types.User
	json.NewDecoder(resp.Body).Decode(&user)

	if user.FirstName != params.FirstName {
		t.Errorf("expected firstname %s but got %s", user.FirstName, params.FirstName)
	}
	if user.LastName != params.LastName {
		t.Errorf("expected last name %s but got %s", user.LastName, params.LastName)
	}
	if user.Email != params.Email {
		t.Errorf("expected email %s but got %s", user.Email, params.Email)
	}
	if len(user.EncryptedPassword) > 0 {
		t.Errorf("ecrypted password must not included")
	}
	if len(user.ID) == 0 {
		t.Errorf("expecting user id set in database")
	}

}
