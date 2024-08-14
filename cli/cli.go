package cli

import (
	"bufio"
	"fmt"
	"os"
)

func RequestInput(request string) (*string, error) {
	fmt.Printf("%s: ", request)
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	response := scanner.Text()

	if err := scanner.Err(); err != nil {
		return nil, err
	}

	return &response, nil
}
