package main

import (
	"net/http"

	"github.com/go-openapi/strfmt"
	"github.com/google/uuid"
	"github.com/labstack/echo/v4"
)

// User for this application
//
// A user is the security principal for this application.
// It's also used as one of main axes for reporting.
//
// swagger:model user
type User struct {
	// the uuid for this user
	//
	// required: true
	Uuid uuid.UUID `json:"uuid"`
	// the id for this user
	//
	// required: true
	// min: 1
	Id int `json:"id"`
	// the name for this user
	//
	// required: true
	// min length: 3
	Name string `json:"name"`
	// the email address for this user
	//
	// required: true
	// example: user@provider.net
	Email strfmt.Email `json:"email"`
	// the age of this user
	//
	// min: 1
	Age int `json:"age"`
}

// Parameters used for creating a new user
//
// swagger:parameters createUser
type CreateUserParameters struct {
	// Name of the created user
	//
	// in: body
	// required: true
	Name string `json:"name"`
	// Email of the created user
	//
	// in: body
	// required: true
	Email string `json:"email"`
}

// User returned on successful creation
//
// swagger:response createUserResponse
type CreateUserResponse struct {
	// in: body
	User User `json:"user"`
}

// swagger:route POST /users createUser
//
// Create user based on required parameters.
//
// This will create an user for application.
// Required are email and password for user creation.
//
// Responses:
//
//	  default: error
//		 200: createUserResponse
func createUser(c echo.Context) error {
	lock.Lock()
	defer lock.Unlock()

	user := &User{}
	if err := c.Bind(user); err != nil {
		return err
	}
	user.Id = seq
	user.Uuid = uuid.New()
	users[user.Id] = user

	seq += 1

	return c.JSON(http.StatusCreated, user)
}

// swagger:route GET /users getAllUsers
//
// Get all users created in application.
//
// This endpoint is responsible for fetching all users.
func getAllUsers(c echo.Context) error {
	lock.Lock()
	defer lock.Unlock()

	return c.JSON(http.StatusOK, users)
}
