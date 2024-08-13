package main

import (
	"fmt"

	"github.com/AidansCode/go-test/model"
)

func main() {
	aidan := model.Person{Name: "Aidan", Age: 99, FavoriteNumber: 6}
	fmt.Println(aidan)
}
