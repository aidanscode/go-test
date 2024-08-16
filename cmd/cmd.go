package cmd

import (
	"errors"
	"strconv"

	"github.com/AidansCode/go-test/model"
)

var commandMappings = map[string] func(args []string, t []model.Task)([]model.Task, bool) {
	"add": Add,
	"list": List,
	"complete": Complete,
	"delete": Delete,
	"sort": Sort,
	"help": Help,
}

const defaultCommand string = "list"

func ExecuteCommand(args []string, tasks []model.Task) ([]model.Task, bool, error) {
	command, subArgs := getCommand(args)
	handler, exists := commandMappings[command]
	if !exists {
		return tasks, false, errors.New("Command does not exist: " + command)
	}

	tasks, isDirty := handler(subArgs, tasks)
	return tasks, isDirty, nil
}

func getCommand(args []string) (string, []string) {
	if len(args) > 1 {
		return args[1], args[2:]
	} else {
		return defaultCommand, make([]string, 0)
	}
}

func findTaskIdFromArgs(args []string) (int, error) {
	if len(args) == 0 {
		return 0, errors.New("no task id provided")
	}

	id, err := strconv.Atoi(args[0])
	if err != nil {
		return 0, errors.New("invalid task id provided")
	}

	return id, nil
}
