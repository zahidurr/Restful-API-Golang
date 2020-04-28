package handlers

import (
	"net/http"

	"github.com/zahidurr/Restful-API-Golang/data"
)

// swagger:route POST /users users createUser
// Create a new user
//
// responses:
//	200: userResponse
//  422: errorValidation
//  501: errorResponse

// Create handles POST requests to add new users
func (p *Users) Create(rw http.ResponseWriter, r *http.Request) {
	// fetch the user from the context
	prod := r.Context().Value(KeyUser{}).(data.User)

	p.l.Printf("[DEBUG] Inserting user: %#v\n", prod)
	data.AddUser(prod)
}
