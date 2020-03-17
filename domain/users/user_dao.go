package users

import (
	"fmt"

	"github.com/sachinkapalidigi/bookstore_users-api/utils/errors"
)

var (
	userDB = make(map[int64]*User)
)

func (user *User) Get() *errors.RestErr {

	result := userDB[user.ID]
	if result == nil {
		return errors.NewNotFoundError(fmt.Sprintf("User %d not found", user.ID))
	}
	user.ID = result.ID
	user.FirstName = result.FirstName
	user.LastName = result.LastName
	user.Email = result.Email
	user.DateCreated = result.DateCreated
	return nil
}

func (user *User) Save() *errors.RestErr {
	current := userDB[user.ID]
	if userDB[user.ID] != nil {
		if current.Email == user.Email {
			return errors.NewBadRequestError("Email already exists")
		}
		return errors.NewBadRequestError("User already exists")
	}

	return nil
}
