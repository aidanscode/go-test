package main

import (
	"os"

	"github.com/AidansCode/go-test/cmd"
	"github.com/AidansCode/go-test/fileio"
)

func main() {
	tasks := fileio.GetTasks()
	tasks, isDirty, err := cmd.ExecuteCommand(os.Args, tasks)
	if err != nil {
		panic(err)
	}
	if isDirty {
		fileio.SaveTasks(tasks)
	}
}
