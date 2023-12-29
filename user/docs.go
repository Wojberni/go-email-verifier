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

// swagger:route GET /users/{id} GetUser
//
// Get user based on id.
//
// This will find and return an user for application given id.
//
// Responses:
//
//	  default: error
//		 201: user

// swagger:route PUT /users/{id} UpdateUser
//
// Update user given id based on required parameters.
//
// This will update an user for application.
// Required are email and password for user creation.
//
// Responses:
//
//	  default: error
//		 201: user

// swagger:route DELETE /users/{id} DeleteUser
//
// Delete user based on id.
//
// This will delete an user for application.
//
// Responses:
//
//	  default: error
//		 201: user
