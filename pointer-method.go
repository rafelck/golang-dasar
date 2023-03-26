package main

import "fmt"

type Man struct {
	Name string
}

func (man *Man) Married() {
	man.Name = "Mr." + man.Name
}

func main() {
	var person Man
	person = Man{
		Name: "Rafel",
	}

	person.Married()

	fmt.Println(person)
}
