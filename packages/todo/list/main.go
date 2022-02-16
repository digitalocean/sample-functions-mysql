package main

import (
	"github.com/jonfriesen/todo"
)

func Main(args map[string]interface{}) map[string]interface{} {
	res := make(map[string]interface{})

	includeCompleted := false
	if i, ok := args["includeCompleted"].(bool); ok && i {
		includeCompleted = i
	}

	todos := todo.List(includeCompleted)

	res["body"] = todos

	return res
}
