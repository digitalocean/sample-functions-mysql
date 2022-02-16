package main

import (
	"fmt"

	"github.com/iancoleman/strcase"
)

func Main(args map[string]interface{}) map[string]interface{} {
	res := make(map[string]interface{})
	greetings := "world"
	name, ok := args["name"].(string)
	if ok {
		greetings = name
	}

	screamingSnake := strcase.ToScreamingSnake(greetings)

	res["package-main"] = "Hello, " + screamingSnake

	fmt.Printf("Hello, %s\n", screamingSnake)
	return res
}
