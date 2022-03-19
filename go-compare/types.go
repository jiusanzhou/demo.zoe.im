package main

import "fmt"

type Person struct {
	Name     string
	Age      int
	Gender   int
	Email    string
	Children []Person
}

func (p Person) String() string {
	return fmt.Sprintf("%s (%d)", p.Name, p.Age)
}
