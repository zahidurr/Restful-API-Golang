package handlers

import (
	"net/http"

	"github.com/zahidurr/Restful-API-Golang/data"
)

// swagger:route PUT /users users updateUser
// Update a users details
//
// responses:
//	201: noContentResponse
//  404: errorResponse
//  422: errorValidation

// Update handles PUT requests to update users
func (p *Users) Update(rw http.ResponseWriter, r *http.Request) {
	rw.Header().Add("Content-Type", "application/json")

	// fetch the user from the context
	prod := r.Context().Value(KeyUser{}).(data.User)
	p.l.Println("[DEBUG] updating record id", prod.ID)

	err := data.UpdateUser(prod)
	if err == data.ErrUserNotFound {
		p.l.Println("[ERROR] user not found", err)

		rw.WriteHeader(http.StatusNotFound)
		data.ToJSON(&GenericError{Message: "User not found in database"}, rw)
		return
	}

	// write the no content success header
	rw.WriteHeader(http.StatusNoContent)
}
