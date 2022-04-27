package main

import (
	"context"
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
	ctx := context.Background()
	// Create a table.
	_, err = db.ExecContext(ctx, `CREATE TABLE IF NOT EXISTS hello(
                                         id int(11) AUTO_INCREMENT PRIMARY KEY,
                                         value int(11)
                                      )`)
	if err != nil {
		res["body"] = err.Error()
		return res
	}
	// Insert 10 rows in db
	for i := 0; i < 10; i++ {
		stmt, err := db.PrepareContext(ctx, "INSERT INTO hello(value) VALUES(?)")
		if err != nil {
			res["body"] = err.Error()
			return res
		}
		_, err = stmt.ExecContext(ctx, i)
		if err != nil {
			res["body"] = err.Error()
			return res
		}
	}
	// Select max value inserted.
	r := db.QueryRow("SELECT MAX(id) FROM hello")
	if err != nil {
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
