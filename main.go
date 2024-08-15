package main

import (
	"fmt"

	"github.com/AidansCode/go-test/fileio"
)

func main() {
	tasks := fileio.GetTasks()
	fmt.Println("Received list", tasks)
	// isDirty, err := cmd.ExecuteCommand(os.Args, tasks)
	// if err != nil {
	// 	panic(err)
	// }
	// if isDirty {
	// 	fileio.SaveTasks(tasks)
	// }
}
