// Package handlers classification of User API
//
// Documentation for User API
//
//	Schemes: http
//	BasePath: /
//	Version: 1.0.0
//
//	Consumes:
//	- application/json
//
//	Produces:
//	- application/json
//
// swagger:meta
package handlers

import "github.com/zahidurr/Restful-API-Golang/data"

//
// NOTE: Types defined here are purely for documentation purposes
// these types are not used by any of the handers

// Generic error message returned as a string
// swagger:response errorResponse
type errorResponseWrapper struct {
	// Description of the error
	// in: body
	Body GenericError
}

// Validation errors defined as an array of strings
// swagger:response errorValidation
type errorValidationWrapper struct {
	// Collection of the errors
	// in: body
	Body ValidationError
}

// A list of users
// swagger:response usersResponse
type usersResponseWrapper struct {
	// All current users
	// in: body
	Body []data.User
}

// Data structure representing a single user
// swagger:response userResponse
type userResponseWrapper struct {
	// Newly created user
	// in: body
	Body data.User
}

// No content is returned by this API endpoint
// swagger:response noContentResponse
type noContentResponseWrapper struct {
}

// swagger:parameters updateUser createUser
type userParamsWrapper struct {
	// User data structure to Update or Create.
	// Note: the id field is ignored by update and create operations
	// in: body
	// required: true
	Body data.User
}

// swagger:parameters listSingleUser deleteUser
type userIDParamsWrapper struct {
	// The id of the user for which the operation relates
	// in: path
	// required: true
	ID int `json:"id"`
}
