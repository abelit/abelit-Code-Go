package main

import (
	"database/sql"
	"fmt"

	_ "gopkg.in/goracle.v2"
)

func main() {
	db, _ := sql.Open("goracle", "system/Passw0rd@10.10.10.204/db19c.dataforum.org")
	defer db.Close()
	db.Ping()
	rows, err := db.Query("select sysdate from dual")
	if err != nil {
		fmt.Println("Error running query")
		fmt.Println(err)
		return
	}
	defer rows.Close()

	var thedate string
	for rows.Next() {

		rows.Scan(&thedate)
	}
	fmt.Printf("The date is: %s\n", thedate)
}
