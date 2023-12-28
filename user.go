package main

import (
	"github.com/go-openapi/strfmt"
	"github.com/google/uuid"
	"github.com/labstack/echo/v4"
)

// User represents the user for this application
//
// A user is the security principal for this application.
// It's also used as one of main axes for reporting.
//
// A user can have friends with whom they can share what they like.
//
// swagger:model user
type User struct {
	Uuid uuid.UUID `json:"uuid"`
	// the id for this user
	//
	// required: true
	// min: 1
	Id int `json:"id"`
	// the name for this user
	// required: true
	// min length: 3
	Name string `json:"name"`
	// the email address for this user
	//
	// required: true
	// example: user@provider.net
	Email strfmt.Email `json:"email"`
	// the friends for this user
	//
	// Extensions:
	// ---
	// x-property-value: value
	// x-property-array:
	//   - value1
	//   - value2
	// x-property-array-obj:
	//   - name: obj
	//     value: field
	// ---
	Age int `json:"age"`
}

func createUser(c echo.Context) error {
	// swagger:route POST /users createUser
	//
	// Lists pets filtered by some parameters.
	//
	// This will show all available pets by default.
	// You can get the pets that are out of stock
	//
	//
	//     Security:
	//       api_key:
	//       oauth: read, write
	//
	//     Parameters:
	//       + name: limit
	//         in: query
	//         description: maximum numnber of results to return
	//         required: false
	//         type: integer
	//         format: int32
	//
	//
	//     Responses:
	//       default: error
	//       200: user
	//       422: error

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

	return nil
}
