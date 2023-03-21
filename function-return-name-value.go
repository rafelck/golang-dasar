package main

import "fmt"

func getFullName2() (firstName, middleName, lastName string) {
	firstName = "Rafelino"
	middleName = "Claudius"
	lastName = "Kelen"

	return
}

func main() {
	a, b, c := getFullName2()
	fmt.Println(a, b, c)
}
