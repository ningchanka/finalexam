package customer

import (
	"database/sql"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/ningchanka/finalexam/database"
	"github.com/ningchanka/finalexam/errors"
)

func UpdateCustomerHandler(c *gin.Context) {
	var customer Customer
	if err := c.ShouldBindJSON(&customer); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	err := updateCustomerService(database.Conn(), &customer)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, customer)
}

func updateCustomerService(db *sql.DB, customer *Customer) error {
	stmt, err := db.Prepare("UPDATE customer SET name=$2, email=$3, status=$4 WHERE id = $1")
	if err != nil {
		return errors.New(err, 666, "can't prepare update statement")
	}

	_, err = stmt.Exec(customer.ID, customer.Name, customer.Email, customer.Status)
	if err != nil {
		return errors.New(err, 666, "can't exec update")
	}

	return nil
}
