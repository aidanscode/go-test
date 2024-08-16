package cmd

import (
	"errors"
	"strings"

	"github.com/AidansCode/go-test/model"
)

func Add(args []string, tasks []model.Task) ([]model.Task, bool) {
	taskName, err := getNewTaskName(args)
	if err != nil {
		panic(err)
	}

	tasks = append(tasks, model.Task{Name: *taskName, Complete: false})
	return tasks, true
}

func getNewTaskName(args []string) (*string, error) {
	if len(args) == 0 {
		return nil, errors.New("no task name provided")
	}

	taskName := strings.Join(args, " ")
	return &taskName, nil
}
