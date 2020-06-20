package customer

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/ningchanka/finalexam/database"
)

func CreateCustomerHandler(c *gin.Context) {
	var customer Customer
	if err := c.ShouldBindJSON(&customer); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	stmt, err := database.Conn().Prepare("INSERT INTO customer (name, email, status) VALUES ($1, $2, $3) RETURNING id")
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	row := stmt.QueryRow(customer.Name, customer.Email, customer.Status)
	var id int
	err = row.Scan(&id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	customer.ID = id

	c.JSON(http.StatusCreated, customer)
}
