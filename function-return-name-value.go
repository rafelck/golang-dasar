package main

import "fmt"

func getFullName() (firstName, middleName, lastName string) {
	firstName = "Rafelino"
	middleName = "Claudius"
	lastName = "Kelen"

	return
}

func main() {
	a, b, c := getFullName()
	fmt.Println(a, b, c)
}
