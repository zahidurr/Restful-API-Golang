package data

import "fmt"

// ErrUserNotFound is an error raised when a user can not be found in the database
var ErrUserNotFound = fmt.Errorf("User not found")

// User defines the structure for an API user
// swagger:model
type User struct {
	// the id for the user
	//
	// required: false
	// min: 1
	ID int `json:"id"` // Unique identifier for the user

	// the First name of this user
	//
	// required: true
	// max length: 15
	Firstname string `json:"first_name"`

	// the Last name of this user
	//
	// required: true
	// max length: 15
	Lastname string `json:"last_name"`

	// the date of birth of this user
	//
	// required: true
	DOB string `json:"dob"`

	// the address of this user
	//
	// required: false
	Address string `json:"address"`
}

// Users defines a slice of User
type Users []*User

// GetUsers returns all users from the database
func GetUsers() Users {
	return userList
}

// GetUserByID returns a single user which matches the id from the
// database.
// If a user is not found this function returns a UserNotFound error
func GetUserByID(id int) (*User, error) {
	i := findIndexByUserID(id)
	if id == -1 {
		return nil, ErrUserNotFound
	}

	return userList[i], nil
}

// UpdateUser replaces a user in the database with the given
// item.
// If a user with the given id does not exist in the database
// this function returns a UserNotFound error
func UpdateUser(p User) error {
	i := findIndexByUserID(p.ID)
	if i == -1 {
		return ErrUserNotFound
	}

	// update the user in the DB
	userList[i] = &p

	return nil
}

// AddUser adds a new user to the database
func AddUser(p User) {
	// get the next id in sequence
	maxID := userList[len(userList)-1].ID
	p.ID = maxID + 1
	userList = append(userList, &p)
}

// DeleteUser deletes a user from the database
func DeleteUser(id int) error {
	i := findIndexByUserID(id)
	if i == -1 {
		return ErrUserNotFound
	}

	userList = append(userList[:i], userList[i+1])

	return nil
}

// findIndex finds the index of a user in the database
// returns -1 when no user can be found
func findIndexByUserID(id int) int {
	for i, p := range userList {
		if p.ID == id {
			return i
		}
	}

	return -1
}

var userList = []*User{
	{
		ID:        1,
		Firstname: "James",
		Lastname:  "bond",
		DOB:       "1991-10-11",
		Address:   "Sweden",
	},
	{
		ID:        2,
		Firstname: "John ",
		Lastname:  "Cena",
		DOB:       "1992-11-11",
		Address:   "USA",
	},
}
