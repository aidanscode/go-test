package cmd

import (
	"github.com/AidansCode/go-test/model"
)

func Delete(args []string, tasks []model.Task) ([]model.Task, bool) {
	taskId, err := findTaskIdFromArgs(args)
	if err != nil {
		panic(err)
	}

	index := taskId - 1
	tasks = append(tasks[:index], tasks[index+1:]...)

	return tasks, true
}
