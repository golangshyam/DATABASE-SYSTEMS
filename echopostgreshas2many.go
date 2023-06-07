package main

import (
	"database/sql"
	"fmt"
	"log"
	"net/http"

	"github.com/labstack/echo"
	_ "github.com/lib/pq"
)

type Employee struct {
	Id     string `json:"id"`
	Name   string `json:"name"`
	Salary string `json: "salary"`
	Age    string `json : "age"`
}
type Employees struct {
	Employees []Employee `json:"employees"`
}

var err error

func main() {
	db, err = sql.Open("postgres", "user=postgres password=root dbname=books_database sslmode=disable")
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	if err = db.Ping(); err != nil {
		panic(err)
	} else {
		fmt.Println("DB Connected...")
	}

	e := echo.New()

	e.POST("/employee", func(c echo.Context) error {
		u := new(Employee)
		if err := c.Bind(u); err != nil {
			return err
		}
		sqlStatement := "INSERT INTO employees (name, salry,age)VALUES ($1, $2, $3)"
		res, err := db.Query(sqlStatement, u.Name, u.Salary, u.Age)
		if err != nil {
			fmt.Println(err)
		} else {
			fmt.Println(res)
			return c.JSON(http.StatusCreated, u)
		}
		return c.String(http.StatusOK, "ok")
	})
}
