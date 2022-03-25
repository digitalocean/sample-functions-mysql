package main

import (
	"database/sql"
	"fmt"
	"os"

	_ "github.com/go-sql-driver/mysql"
)

// Main is the entry point of function.
func Main(args map[string]interface{}) map[string]interface{} {
	res := make(map[string]interface{})
	dbURL := os.Getenv("DATABASE_URL")
	if dbURL == "" {
		res["body"] = "invalid db url"
		return res
	}
	dbPassword := os.Getenv("DATABASE_PASSWORD")
	if dbPassword == "" {
		res["body"] = "invalid db password"
		return res
	}
	dsn := fmt.Sprintf("doadmin:%v@tcp(%v:25060)/defaultdb", dbPassword, dbURL)
	db, err := sql.Open("mysql", dsn)
	defer db.Close()
	if err != nil {
		fmt.Println("could not connect to database")
		res["body"] = err.Error()
		return res
	}
	r := db.QueryRow("SELECT 1 FROM DUAL")
	if err != nil {
		fmt.Println("could not query from db")
		res["body"] = err.Error()
		return res
	}
	var result uint32
	err = r.Scan(&result)
	if err != nil {
		res["body"] = err.Error()
		return res
	}
	res["body"] = result
	return res
}
