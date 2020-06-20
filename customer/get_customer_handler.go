package customer

import (
	"database/sql"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/ningchanka/finalexam/database"
	"github.com/ningchanka/finalexam/errors"
)

func GetCustomerByIdHandler(c *gin.Context) {
	id := c.Param("id")
	
	customer, err := getCustomerByIdService(database.Conn(), id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, customer)
}

func getCustomerByIdService(db *sql.DB, id string) (*Customer, error) {
	stmt, err := db.Prepare("SELECT id, name, email, status FROM customer WHERE id = $1")
	if err != nil {
		return nil, errors.New(err, 666, "can't prepare select by id statement")
	}

	row := stmt.QueryRow(id)

	cus := &Customer{}
	err = row.Scan(&cus.ID, &cus.Name, &cus.Email, &cus.Status)
	if err != nil {
		return nil, errors.New(err, 666, "can't exec select by id")
	}

	return cus, nil
}