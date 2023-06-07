package main

import (
	"database/sql"
	"fmt"

	_ "github.com/lib/pq"
)

const (
	host     = "localhost"
	port     = 5432
	user     = "postgres"
	password = "shyamvarma"
	dbname   = "Trainings"
)

func main() {

	psqlconn := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable", host, port, user, password, dbname)

	db, err := sql.Open("postgres", psqlconn)

	CheckErrors(err)

	defer db.Close()

	// insert
	// hardcoded

	insertStmt := `insert into "Students"("Name", "Roll_Number") values('shyam', 444)`
	_, e := db.Exec(insertStmt)
	CheckErrors(e)

	// dynamic
	insertDynStmt := `insert into "Students"("Name", "Roll_Number") values($1, $2)`
	_, e = db.Exec(insertDynStmt, "shyamvarma", 4)
	CheckErrors(e)
}

func CheckErrors(err error) {
	if err != nil {
		panic(err)
	}
}
