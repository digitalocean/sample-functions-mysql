package main

import (
	"fmt"
	"os"
	"strings"

	_ "github.com/go-sql-driver/mysql"
	"github.com/xo/dburl"
)

// Main is the entry point of function.
func Main(args map[string]interface{}) map[string]interface{} {
	res := make(map[string]interface{})
	db, err := dburl.Open(strings.TrimSuffix(os.Getenv("DATABASE_URL"), "?ssl-mode=REQUIRED"))
	if err != nil {
		res["body"] = err.Error()
		return res
	}
	defer db.Close()
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
