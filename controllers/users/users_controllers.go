package users

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/sachinkapalidigi/bookstore_users-api/domain/users"
	"github.com/sachinkapalidigi/bookstore_users-api/services"
	"github.com/sachinkapalidigi/bookstore_users-api/utils/errors"
)

// CreateUser : create a new user
func CreateUser(c *gin.Context) {
	var user users.User
	// bytes, err := ioutil.ReadAll(c.Request.Body)

	// fmt.Println(string(bytes))

	// if err != nil {
	// 	// handle error
	// 	return
	// }

	// if err := json.Unmarshal(bytes, &user); err != nil {
	// 	fmt.Println(err.Error())

	// 	// handle json error
	// 	return
	// }
	// fmt.Println(user)
	if err := c.ShouldBindJSON(&user); err != nil {
		// handle json error
		fmt.Println(err.Error())
		restErr := errors.NewBadRequestError("Invalid JSON body")
		c.JSON(restErr.Status, restErr)
		return
	}

	result, saveError := services.CreateUser(user)

	if saveError != nil {
		// handle save error
		c.JSON(saveError.Status, saveError)
		return
	}

	c.JSON(http.StatusCreated, result)
	// c.String(http.StatusNotImplemented, "implement me!")
}

// GetUser : get a new user
func GetUser(c *gin.Context) {
	userId, userErr := strconv.ParseInt(c.Param("user_id"), 10, 64)
	if userErr != nil {
		err := errors.NewBadRequestError("User Id should be a number")
		c.JSON(err.Status, err)
		return
	}
	user, saveError := service.GetUser(userId)

	if saveError != nil {
		c.JSON(saveErr.Status, saveError)
		return
	}

	c.JSON(http.StatusOK, user)
}
