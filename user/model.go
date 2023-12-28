package user

import (
	"github.com/go-openapi/strfmt"
	"github.com/google/uuid"
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
