package cmd

import (
	"errors"
	"fmt"

	"github.com/AidansCode/go-test/model"
)

var commandMappings = map[string] func(t []model.Task)bool {
	"add": add,
	"list": list,
	"complete": complete,
	"delete": delete,
	"sort": sort,
}

func ExecuteCommand(args []string, tasks []model.Task) (bool, error) {
	command := getCommand(args)
	handler, exists := commandMappings[command]
	if !exists {
		return false, errors.New("Command does not exist: " + command)
	}

	return handler(tasks), nil
}

func getCommand(args []string) string {
	if len(args) > 1 {
		return args[1]
	} else {
		return "list"
	}
}

func add(tasks []model.Task) bool {
	fmt.Println("Adding task")
	return true
}

func list(tasks []model.Task) bool {
	fmt.Println("Listing tasks")
	return false
}

func complete(tasks []model.Task) bool {
	fmt.Println("Completing task")
	return true
}

func delete(tasks []model.Task) bool {
	fmt.Println("Deleting task")
	return true
}

func sort(tasks []model.Task) bool {
	fmt.Println("Sorting tasks")
	return true
}
