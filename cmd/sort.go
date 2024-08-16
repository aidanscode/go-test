package cmd

import (
	"github.com/AidansCode/go-test/model"
)

func Sort(args []string, tasks []model.Task) ([]model.Task, bool) {
	shouldReverseOrder := shouldReverseOrder(args)

	incompleteTasks := filterByIsComplete(tasks, false)
	completeTasks := filterByIsComplete(tasks, true)

	if !shouldReverseOrder {
		tasks = append(incompleteTasks, completeTasks...)
	} else {
		tasks = append(completeTasks, incompleteTasks...)
	}

	return tasks, true
}

func shouldReverseOrder(args []string) bool {
	return len(args) > 0 && args[0] == "reverse"
}

func filterByIsComplete(tasks []model.Task, isComplete bool) []model.Task {
	filteredTasks := make([]model.Task, 0)
	for _, task := range tasks {
		if task.Complete == isComplete {
			filteredTasks = append(filteredTasks, task)
		}
	}

	return filteredTasks
}
