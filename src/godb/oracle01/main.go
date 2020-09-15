package main

import (
	"database/sql"
	"fmt"

	_ "github.com/godror/godror"
)

/**
1. 安装oracle client

2. 配置库依赖
sudo sh -c "echo /usr/lib/oracle/18.3/client64/lib > /etc/ld.so.conf.d/oracle-instantclient.conf"
sudo ldconfig

3. 安装godror driver
go get github.com/godror/godror
**/

func main() {

	db, err := sql.Open("godror", "system/Passw0rd@10.10.10.204/db19c.dataforum.org")
	if err != nil {
		fmt.Println(err)
		return
	}
	defer db.Close()

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

	cs := GetChar()
	fmt.Printf("The CharaterSet is: %s\n", cs)
}

// GetChar is a function
func GetChar() string {
	db, err := sql.Open("godror", "system/Passw0rd@10.10.10.204/db19c.dataforum.org")
	if err != nil {
		fmt.Println(err)
		return "Err"
	}
	defer db.Close()

	rows, err := db.Query("select userenv('LANGUAGE') from dual")
	if err != nil {
		fmt.Println("Error running query")
		fmt.Println(err)
		return "Err"
	}
	defer rows.Close()

	var charset string
	for rows.Next() {

		rows.Scan(&charset)
	}

	return charset
}
