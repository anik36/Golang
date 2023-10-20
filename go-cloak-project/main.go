package main

import (
	"example/goframework/webservice/token"
	"example/goframework/webservice/users"

	"github.com/gin-gonic/gin"
)

func main() {
	// Setup gin router
	router := gin.Default()

	// Token api handlers
	token.RegisterTokenHandlers(router)

	// Register api handlers
	users.RegisterUserHandlers(router)

	router.Run(":8081")

}
