package customer

import (
	"database/sql"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/ningchanka/finalexam/database"
	"github.com/ningchanka/finalexam/errors"
)

func CreateCustomerHandler(c *gin.Context) {
	var customer Customer
	if err := c.ShouldBindJSON(&customer); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	err := createCustomerService(database.Conn(), &customer)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, customer)
}

func createCustomerService(db *sql.DB, customer *Customer) error {
	stmt, err := db.Prepare("INSERT INTO customer (name, email, status) VALUES ($1, $2, $3) RETURNING id")
	if err != nil {
		return errors.New(err, 500, 1000, "can't prepare insert statement")
	}

	row := stmt.QueryRow(customer.Name, customer.Email, customer.Status)
	var id int
	err = row.Scan(&id)
	if err != nil {
		return errors.New(err, 500, 1001, "can't exec insert")
	}
	customer.ID = id
	return nil
}
