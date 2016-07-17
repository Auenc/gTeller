package filter

import "github.com/auenc/gTeller/users"

type UserFilter struct {
	UUID     Condition
	Username Condition
	Email    Condition
	Password Condition
}

//Valid returns true if the given user matches the filter
func (filter *UserFilter) Valid(user users.User) bool {
	valid := true
	emptyCon := Condition{}

	if filter.UUID != emptyCon {
		//Checking if user matches UUID filter.
		//IF it doesn't, return false.
		if !filter.UUID.Valid(user.UUID) {
			valid = false
			return valid
		}
	}

	if filter.Username != emptyCon {
		//Checking if user matches Username filter.
		//IF it doesn't, return false.
		if !filter.Username.Valid(user.Username) {
			valid = false
			return valid
		}
	}

	if filter.Email != emptyCon {
		//Checking if user matches Email filter.
		//IF it doesn't, return false.
		if !filter.Email.Valid(user.Email) {
			valid = false
			return valid
		}
	}

	return valid
}

//Filter filters out Users that do not match the filter from a given slice
func (filter *UserFilter) Filter(source []users.User) []users.User {
	filtered := make([]users.User, 0)

	//Lopping through users
	for _, user := range source {
		if filter.Valid(user) {
			filtered = append(filtered, user)
		}
	}

	return filtered
}
