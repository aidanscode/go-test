package model

import "fmt"

const checkmark string = "✓"

type Task struct {
	Name string
	Complete bool
}

func (task Task) String() string {
	if task.Complete {
		return fmt.Sprintf("%s %s", checkmark, task.Name)
	} else {
		return task.Name
	}
}
