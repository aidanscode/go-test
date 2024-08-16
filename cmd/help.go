package cmd

import (
	"fmt"

	"github.com/AidansCode/go-test/model"
)

func Help(args []string, tasks []model.Task) ([]model.Task, bool) {
	helpText := "Tasky Help:\n" +
		"\ttasky help - displays this help menu\n" +
		"\ttasky - same as 'tasky list'\n" +
		"\ttasky list - lists all tasks\n" +
		"\ttasky add <name> - adds a new task with given name\n" +
		"\ttasky delete <id> - deletes the task with the given id\n" +
		"\ttasky complete <id> - marks the task of the given id as complete\n" +
		"\ttasky sort [reverse] - sorts the task list by completeness, either first or last depending on reverse flag\n"
		fmt.Println(helpText)
	return tasks, false
}
