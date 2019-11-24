package main

import (
	"fmt"

	"github.com/roger-king/tasker/utils"
)

type payload struct {
	Name string `json:"name" validate:"required"`
}

// Step 1: create function called "Run"
func Run(args map[string]interface{}) error {
	var p payload

	// Step 2: call tasker validate function
	// This will ensure we are setting the proper payload for the task
	err := utils.Validate(args, p)

	if err != nil {
		return err
	}

	fmt.Println("Hello Tasker: ", p.Name)
	return nil
}
