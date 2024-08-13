package model

import "fmt"

type Person struct {
	Name string
	Age uint
	FavoriteNumber int
}

func (p Person) String() string {
	return fmt.Sprintf("%s, age %d, has favorite number %d", p.Name, p.Age, p.FavoriteNumber)
}
