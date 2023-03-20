package main

import "fmt"

func main() {
	const company string = "Bank Akp"

	var (
		name   string
		age    int8
		status bool
	)

	name = "Rafelino Claudius Kelen"
	fmt.Println(name)

	name = "Rafelino Kelen"
	fmt.Println(name)

	age = 26
	fmt.Println(age)

	fmt.Println(company)

	status = true
	fmt.Println("Status Pacaran", status)
}
