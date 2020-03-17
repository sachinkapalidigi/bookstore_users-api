package app

import (
	"github.com/gin-gonic/gin"
)

// private method: available only in this package
var (
	router = gin.Default()
)

// StartApplication : start the application
func StartApplication() {
	mapUrls()
	router.Run(":8080")
}
