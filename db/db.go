package db

import (
	"database/sql"
	errors "web-service-gin/error-handle"
	"web-service-gin/repository"
)

var mysqlDB *sql.DB = nil

func connectDB() *sql.DB {
	config := getConfig()

	var err error

	mysqlDB, err = sql.Open("mysql", config.FormatDSN())

	errors.Fatal(err)

	pingErr := mysqlDB.Ping()

	errors.Fatal(pingErr)

	return mysqlDB
}

func GetEmployees(limit uint64) []repository.Employee {
	var employees []repository.Employee

	mysqlDB := getDB()

	rows, err := mysqlDB.Query("SELECT id, title, post FROM employee LIMIT ?", limit)

	errors.Fatal(err)
	defer closeRows(rows)

	for rows.Next() {
		var employee repository.Employee

		err := rows.Scan(&employee.ID, &employee.Name, &employee.Post)
		errors.Fatal(err)

		employees = append(employees, employee)
	}

	errRow := rows.Err()
	errors.Fatal(errRow)

	return employees
}

func GetEmployeeByID(id uint64) (repository.Employee, error) {

	mysqlDB := getDB()
	var employee repository.Employee

	row := mysqlDB.QueryRow("SELECT id, title, post FROM employee WHERE id = ?", id)

	err := row.Scan(&employee.ID, &employee.Name, &employee.Post)
	if err != nil {
		if err == sql.ErrNoRows {
			return employee, sql.ErrNoRows
		}

		errors.Fatal(err)
	}

	return employee, nil
}

func getDB() *sql.DB {
	if mysqlDB == nil {
		return connectDB()
	}

	return mysqlDB
}

func closeRows(rows *sql.Rows) {
	err := rows.Close()
	errors.Fatal(err)
}
