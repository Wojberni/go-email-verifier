package user

// swagger:route GET /users GetAllUsers
//
// Get all users created in application.
//
// This endpoint is responsible for fetching all users.
//
// Responses:
//
//	  default: error
//		 200: []user

// swagger:route POST /users CreateUser
//
// Create user based on required parameters.
//
// This will create an user for application.
// Required are email and password for user creation.
//
// Responses:
//
//	  default: error
//		 201: user
