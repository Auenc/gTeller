package gTeller

import (
	"github.com/auenc/gTeller/filter"
	"github.com/auenc/gTeller/users"
)

//ShippingType objects
type ListUserRequest struct {
	Filter filter.UserFilter
}

type AddUserRequest struct {
	Username string
	Email    string
	Password string
}

type UpdateUserRequest struct {
	Users []users.User
}

type RemoveUserRequest struct {
	Users []users.User
}
