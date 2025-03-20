package main

import (
	"fmt"

	engine "github.com/gabrielbovilorandi/go-template"
)

func main() {
	query, args := engine.Execute(
		"test/app.sql",
		map[string]map[string]string{
			"args": {"id": "123"},
			"data": {"name": "template engine", "version": "1.0.0"},
		},
	)

	fmt.Println(query, args)
}
