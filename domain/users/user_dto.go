package users

import (
	"strings"

	"github.com/sachinkapalidigi/bookstore_users-api/utils/date_utils"
	"github.com/sachinkapalidigi/bookstore_users-api/utils/errors"
)

// User : for user data
type User struct {
	ID          int64  `json:"id"`
	FirstName   string `json:"first_name"`
	LastName    string `json:"last_name"`
	Email       string `json:"email"`
	DateCreated string `json:"date_created"`
}

func (user *User) Validate() *errors.RestErr {
	user.Email = strings.TrimSpace(strings.ToLower(user.Email))
	if user.Email == "" {
		return errors.NewBadRequestError("invalid email address")
	}
	user.DateCreated = date_utils.GetNowString()
	return nil
}
