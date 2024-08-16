package model

import "fmt"

const completeMarker string = "[✓]"

type Task struct {
	Name string
	Complete bool
}

func (task Task) String() string {
	if task.Complete {
		return fmt.Sprintf("%s %s", completeMarker, task.Name)
	} else {
		return task.Name
	}
}
