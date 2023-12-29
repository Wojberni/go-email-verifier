package user

// Wrapper for docs which contains struct used for creating a new user
//
// swagger:parameters CreateUser UpdateUser
type CreateUserDto struct {
	// These are parameters required for creating a new user
	//
	// in: body
	// required: true
	Body struct {
		Name  string `json:"name"`
		Email string `json:"email"`
		Age   int64  `json:"age"`
	} `json:"body"`
}

// An UserID parameter model.
//
// This is used for operations that want the ID of an order in the path
// swagger:parameters GetUser DeleteUser UpdateUser
type UserID struct {
	// The ID of the order
	//
	// in: path
	// required: true
	ID int64 `json:"id"`
}
