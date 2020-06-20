package customer

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func SetupRouter() *gin.Engine {
	r := gin.Default()

	r.Use(AuthMiddleware)

	r.POST("/customers", CreateCustomerHandler)
	r.GET("/customers/:id", tempHandler)
	r.GET("/customers", tempHandler)
	r.PUT("/customers/:id", tempHandler)
	r.DELETE("/customers/:id", tempHandler)
	return r
}

func tempHandler(c *gin.Context) {
	resp := &struct {
		message string
	}{
		message: "temp",
	}
	c.JSON(http.StatusOK, resp)
}
