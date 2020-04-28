package handlers

import (
	"net/http"

	"github.com/zahidurr/Restful-API-Golang/data"
)

// swagger:route GET /users users listUsers
// Return a list of users from the database
// responses:
//	200: usersResponse

// ListAll handles GET requests and returns all current users
func (p *Users) ListAll(rw http.ResponseWriter, r *http.Request) {
	p.l.Println("[DEBUG] get all records")
	rw.Header().Add("Content-Type", "application/json")

	prods := data.GetUsers()

	err := data.ToJSON(prods, rw)
	if err != nil {
		// we should never be here but log the error just incase
		p.l.Println("[ERROR] serializing user", err)
	}
}

// swagger:route GET /users/{id} users listSingleUser
// Return a list of users from the database
// responses:
//	200: userResponse
//	404: errorResponse

// ListSingle handles GET requests
func (p *Users) ListSingle(rw http.ResponseWriter, r *http.Request) {
	rw.Header().Add("Content-Type", "application/json")

	id := getUserID(r)

	p.l.Println("[DEBUG] get record id", id)

	prod, err := data.GetUserByID(id)

	switch err {
	case nil:

	case data.ErrUserNotFound:
		p.l.Println("[ERROR] fetching user", err)

		rw.WriteHeader(http.StatusNotFound)
		data.ToJSON(&GenericError{Message: err.Error()}, rw)
		return
	default:
		p.l.Println("[ERROR] fetching user", err)

		rw.WriteHeader(http.StatusInternalServerError)
		data.ToJSON(&GenericError{Message: err.Error()}, rw)
		return
	}

	err = data.ToJSON(prod, rw)
	if err != nil {
		// we should never be here but log the error just incase
		p.l.Println("[ERROR] serializing user", err)
	}
}
