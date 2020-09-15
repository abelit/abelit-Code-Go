package main

import (
	"database/sql"
	"fmt"

	_ "github.com/go-sql-driver/mysql"
)

func main() {
	db, err := sql.Open("mysql", "root:root@tcp(10.50.0.210:3306)/mis?charset=utf8")
	if err != nil {
		fmt.Println(err)
		return
	}
	defer db.Close()

	rows, err := db.Query("select sum(gw_salesamount) from gw_transsalestotal")

	if err != nil {
		fmt.Println("Error running query")
		fmt.Println(err)
		return
	}
	defer rows.Close()

	fmt.Println(*rows)

	var result string
	for rows.Next() {
		rows.Scan(&result)
	}

	fmt.Printf("The sales amount is: %s\n", result)
}
