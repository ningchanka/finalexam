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
		c.JSON(http.StatusInternalServerError, err)
		return
	}

	c.JSON(http.StatusOK, customer)
}

func GetCustomersHandler(c *gin.Context) {
	customer, err := getCustomersService(database.Conn())
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, customer)
}

func getCustomerByIdService(db *sql.DB, id string) (*Customer, error) {
	stmt, err := db.Prepare("SELECT id, name, email, status FROM customer WHERE id = $1")
	if err != nil {
		return nil, errors.New(err, 500, 1000, "can't prepare select by id statement")
	}

	row := stmt.QueryRow(id)

	cus := &Customer{}
	err = row.Scan(&cus.ID, &cus.Name, &cus.Email, &cus.Status)
	if err != nil && err == sql.ErrNoRows {
		return nil, errors.New(err, 404, 404, "not found")
	}
	if err != nil {
		return nil, errors.New(err, 500, 1001, "can't exec select by id")
	}

	return cus, nil
}

func getCustomersService(db *sql.DB) ([]*Customer, error) {
	stmt, err := db.Prepare("SELECT id, name, email, status FROM customer")
	if err != nil {
		return nil, errors.New(err, 500, 1000, "can't prepare select all statement")
	}

	rows, err := stmt.Query()
	if err != nil {
		return nil, errors.New(err, 500, 1001, "can't exec select all")
	}

	items := []*Customer{}
	for rows.Next() {
		cus := &Customer{}
		err := rows.Scan(&cus.ID, &cus.Name, &cus.Email, &cus.Status)
		if err != nil {
			return nil, err
		}
		items = append(items, cus)
	}

	return items, nil
}
