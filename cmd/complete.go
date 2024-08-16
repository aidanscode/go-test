package cmd

import "github.com/AidansCode/go-test/model"

func Complete(args []string, tasks []model.Task) ([]model.Task, bool) {
	taskId, err := findTaskIdFromArgs(args)
	if err != nil {
		panic(err)
	}
	tasks[taskId - 1].Complete = true
	return tasks, true
}
