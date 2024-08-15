package fileio

import (
	"encoding/json"
	"os"

	"github.com/AidansCode/go-test/model"
)

func CreateNewDataFile() {
	tasks := make([]model.Task, 0)
	SaveTasks(tasks)
}

func SaveTasks(tasks []model.Task) {
	bytes, err := json.Marshal(tasks)
	if err != nil {
		panic(err)
	}

	if err := os.WriteFile(fileName, bytes, 0640); err != nil {
		panic(err)
	}
}
