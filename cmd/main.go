package main

import (
	"database/sql"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
	"web-service-gin/db"
	errors "web-service-gin/error-handle"
)

func main() {
	router := gin.Default()
	router.GET("/employees", getEmployees)
	router.GET("/employees/:id", getEmployeeByID)

	err := router.Run("localhost:8080")
	errors.Fatal(err)
}

func getEmployees(c *gin.Context) {
	employees := db.GetEmployees(3)

	c.IndentedJSON(http.StatusOK, employees)
}

func getEmployeeByID(c *gin.Context) {
	idRaw := c.Param("id")
	id, err := strconv.ParseUint(idRaw, 10, 64)
	errors.Fatal(err)

	employee, err := db.GetEmployeeByID(id)

	if err == sql.ErrNoRows {
		c.IndentedJSON(http.StatusNotFound, gin.H{"message": "employee not found"})
		return
	}

	c.IndentedJSON(http.StatusOK, employee)
}
