package main

import (
	"database/sql"
	"fmt"

	_ "github.com/lib/pq"
)

func main() {
	connStr := "user=postgres password=shyamvarma host=localhost port=5432 dbname=Students sslmode=disable"
	db, err := sql.Open("postgres", connStr)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}
	defer db.Close()

	stmt, err := db.Prepare(`INSERT INTO Class (Name, Rollnumber) VALUES ($1, $2)`)
	//sqlQuery := `INSERT INTO users ("user_id", "username", "password", "email", "gender", "gang") VALUES ($1, $2, $3, $4, $5, $6)`
	//_, err = db.Exec(sqlQuery, 1, "Ross8839", "rocky8839", "ross88399@hotmail.com", "Female", "Greengos")
	//if err != nil {
	//  fmt.Fprintf(os.Stdout, "Failed to query db")
	//  panic(err)
	//}
	//}
	if err != nil {
		fmt.Println("Error:", err)
		return
	}
	defer stmt.Close()

	result, err := stmt.Exec("shyam", "411")
	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	rowsAffected, err := result.RowsAffected()
	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	fmt.Println("Rows affected:", rowsAffected)
}
