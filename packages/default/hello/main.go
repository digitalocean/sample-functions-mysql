package main

import (
	"database/sql"
	"fmt"

	_ "github.com/go-sql-driver/mysql"
)

// Main is the entry point of function.
func Main(args map[string]interface{}) map[string]interface{} {
	res := make(map[string]interface{})
	db, err := sql.Open("mysql", "doadmin:hyMY0iFw2g72R7zX@db-mysql-nyc3-40125-do-user-4646103-0.b.db.ondigitalocean.com:25060/defaultdb")
	if err != nil {
		fmt.Println("could not connect to database")
		res["body"] = err.Error()
		return res
	}
	r, err := db.Query("SELECT 1 FROM DUAL")
	if err != nil {
		fmt.Println("could not query from db")
		res["body"] = err.Error()
		return res
	}
	var result uint32
	r.Scan(&result)
	res["body"] = result
	return res
}
