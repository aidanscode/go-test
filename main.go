package main

import (
	"fmt"
	"strconv"

	"github.com/AidansCode/go-test/cli"
	"github.com/AidansCode/go-test/model"
)

func main() {
	name, err := cli.RequestInput("Enter your name")
	if err != nil {
		fmt.Println("An error occurred reading your name")
		return
	}

	rawAge, err := cli.RequestInput("Enter your age")
	if err != nil {
		fmt.Println("An error occurred reading your age")
		return
	}

	age, err := strconv.ParseUint(*rawAge, 10, 64)
	if err != nil {
		fmt.Println("An error occurred parsing your age")
		return
	}

	rawFavoriteNumber, err := cli.RequestInput("Enter your favorite number")
	if err != nil {
		fmt.Println("An error occurred reading your favorite number")
		return
	}

	favoriteNumber, err := strconv.Atoi(*rawFavoriteNumber)
	if err != nil {
		fmt.Println("An error occurred parsing your favorite number")
		return
	}

	person := model.Person{Name: *name, Age: uint(age), FavoriteNumber: favoriteNumber}
	fmt.Println(person)
}
