package main

import (
	"fmt"

	"github.com/jonfriesen/todo"
)

func Main(args map[string]interface{}) map[string]interface{} {
	res := make(map[string]interface{})
	itemID, ok := args["item"].(string)
	if !ok {
		res["body"] = "no item found"
		return res
	}

	item, err := todo.Get(itemID)
	if err != nil {
		fmt.Println(">>>", err.Error())
		res["body"] = err.Error()
		return res
	}

	res["body"] = item
	return res
}
