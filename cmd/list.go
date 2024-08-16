package cmd

import (
	"fmt"

	"github.com/AidansCode/go-test/model"
)

func List(args []string, tasks []model.Task) ([]model.Task, bool) {
	if len(tasks) > 0 {
		for index, task := range(tasks) {
			fmt.Printf("%d. %v\n", (index + 1), task)
		}
	} else {
		fmt.Println("(No tasks yet)")
	}
	return tasks, false
}
