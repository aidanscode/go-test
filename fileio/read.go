package fileio

import (
	"encoding/json"
	"errors"
	"io/fs"
	"os"

	"github.com/AidansCode/go-test/model"
)

func GetTasks() []model.Task {
	var tasks []model.Task
	rawContents := readDataFile()

	err := json.Unmarshal(rawContents, &tasks)
	if err != nil {
		panic(errors.New("non-parseable save file"))
	}

	return tasks
}

func readDataFile() []byte {
	byteContent, err := os.ReadFile(fileName)
	if err != nil {
		if errors.Is(err, fs.ErrNotExist) {
			CreateNewDataFile()
			return readDataFile()
		} else {
			panic(err)
		}
	}

	return byteContent
}
