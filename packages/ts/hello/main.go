package main

import (
	"os"

	_ "github.com/go-sql-driver/mysql"
)

// Main is the entry point of function.
func Main(args map[string]interface{}) map[string]interface{} {
	res := make(map[string]interface{})
	dbURL := os.Getenv("DATABASE_URL")
	res["body"] = dbURL
	return res
	// db, err := sql.Open("mysql", dbURL)
	// defer db.Close()
	// if err != nil {
	// 	fmt.Println("could not connect to database")
	// 	res["body"] = dbURL
	// 	return res
	// }
	// r := db.QueryRow("SELECT 1 FROM DUAL")
	// if err != nil {
	// 	fmt.Println("could not query from db")
	// 	res["body"] = err.Error()
	// 	return res
	// }
	// var result uint32
	// err = r.Scan(&result)
	// if err != nil {
	// 	res["body"] = err.Error()
	// 	return res
	// }
	// res["body"] = result
	// return res
}
