package user

// Wrapper for docs which contains struct used for creating a new user
//
// swagger:parameters CreateUser
type CreateUserDto struct {
	// These are parameters required for creating a new user
	//
	// in: body
	// required: true
	Body struct {
		Name  string `json:"name"`
		Email string `json:"email"`
	} `json:"body"`
}
