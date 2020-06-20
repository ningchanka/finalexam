package customer

import (
	"github.com/gin-gonic/gin"
)

func SetupRouter() *gin.Engine {
	r := gin.Default()

	r.Use(AuthMiddleware)

	r.POST("/customers", CreateCustomerHandler)
	r.GET("/customers/:id", GetCustomerByIdHandler)
	r.GET("/customers", GetCustomersHandler)
	r.PUT("/customers/:id", UpdateCustomerHandler)
	r.DELETE("/customers/:id", DeleteCustomerHandler)
	return r
}
