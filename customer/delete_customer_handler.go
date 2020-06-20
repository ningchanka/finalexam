package customer

import (
	"database/sql"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/ningchanka/finalexam/database"
	"github.com/ningchanka/finalexam/errors"
)

func DeleteCustomerHandler(c *gin.Context) {
	id := c.Param("id")

	err := deleteCustomerService(database.Conn(), id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "customer deleted",
	})
}

func deleteCustomerService(db *sql.DB, id string) error {
	stmt, err := db.Prepare("DELETE FROM customer WHERE id = $1")
	if err != nil {
		return errors.New(err, 666, "can't prepare delete statement")
	}

	_, err = stmt.Exec(id)
	if err != nil {
		return errors.New(err, 666, "can't exec delete")
	}

	return nil
}
