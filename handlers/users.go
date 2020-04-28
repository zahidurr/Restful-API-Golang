package handlers

import (
	"fmt"
	"log"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
	"github.com/zahidurr/Restful-API-Golang/data"
)

// KeyUser is a key used for the User object in the context
type KeyUser struct{}

// Users handler for getting and updating users
type Users struct {
	l *log.Logger
	v *data.Validation
}

// NewUsers returns a new users handler with the given logger
func NewUsers(l *log.Logger, v *data.Validation) *Users {
	return &Users{l, v}
}

// ErrInvalidUserPath is an error message when the user path is not valid
var ErrInvalidUserPath = fmt.Errorf("Invalid Path, path should be /users/[id]")

// GenericError is a generic error message returned by a server
type GenericError struct {
	Message string `json:"message"`
}

// ValidationError is a collection of validation error messages
type ValidationError struct {
	Messages []string `json:"messages"`
}

// getUserID returns the user ID from the URL
// Panics if cannot convert the id into an integer
// this should never happen as the router ensures that
// this is a valid number
func getUserID(r *http.Request) int {
	// parse the user id from the url
	vars := mux.Vars(r)

	// convert the id into an integer and return
	id, err := strconv.Atoi(vars["id"])
	if err != nil {
		// should never happen
		panic(err)
	}

	return id
}
