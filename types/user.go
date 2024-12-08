package types

import (
	"encoding/base64"
	"fmt"
	"regexp"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"golang.org/x/crypto/bcrypt"
)

const bcryptCost = 12

type User struct {
	ID                primitive.ObjectID `bson:"_id,omitempty" json:"id,omitempty"`
	FirstName         string             `bson:"first_name" json:"firstName"`
	LastName          string             `bson:"last_name" json:"lastName"`
	Email             string             `bson:"email" json:"email"`
	EncryptedPassword string             `bson:"hashed_password" json:"-" `
}

type CreateUserParams struct {
	FirstName string `json:"firstName"`
	LastName  string `json:"lastName"`
	Email     string `json:"email"`
	Password  string `json:"password"`
}

type UpdateUserParams struct {
	FirstName string `json:"firstName"`
	LastName  string `json:"lastName"`
}

func EncodingUserPassword(params CreateUserParams) (*User, error) {
	ecpw, err := bcrypt.GenerateFromPassword([]byte(params.Password), bcryptCost)
	if err != nil {
		return nil, err
	}

	hashedPassword := base64.StdEncoding.EncodeToString(ecpw)
	return &User{
		FirstName:         params.FirstName,
		LastName:          params.LastName,
		Email:             params.Email,
		EncryptedPassword: hashedPassword,
	}, nil
}

const (
	minLenFirstName = 2
	minLenLastName  = 2
	minLenPassword  = 7
)

func (params CreateUserParams) InputValidation() []string {
	var errors = []string{}

	if len(params.FirstName) < minLenFirstName {
		errors = append(errors, fmt.Sprintf("the minimum length of firstname is: %d", minLenFirstName))
	}
	if len(params.LastName) < minLenLastName {
		errors = append(errors, fmt.Sprintf("the minimum length of last name is: %d", minLenLastName))
	}
	if len(params.Password) < minLenPassword {
		errors = append(errors, fmt.Sprintf("the minimum length of password is: %d", minLenPassword))
	}
	if !isEmailValid(params.Email) {
		errors = append(errors, fmt.Sprintf("email is not valid"))
	}
	return errors
}

func isEmailValid(e string) bool {
	emailRegex := regexp.MustCompile(`^[a-z0-9._%+\-]+@[a-z0-9.\-]+\.[a-z]{2,4}$`)
	return emailRegex.MatchString(e)
}

func (p UpdateUserParams) TOBSON() bson.M {
	m := bson.M{}
	if len(p.FirstName) > minLenFirstName {
		m["firstName"] = p.FirstName
	}
	if len(p.LastName) > minLenLastName {
		m["lastName"] = p.LastName
	}
	return m
}
