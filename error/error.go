package api

// Error object
//
// default struct for error handling
//
// swagger:model error
type Error struct {
	// the error code
	//
	// required: true
	// min: 100
	// max: 599
	// example: 500
	Code int64 `json:"code"`
	// the error message
	//
	// required: true
	// example: error message returned by function
	Message string `json:"message"`
}
