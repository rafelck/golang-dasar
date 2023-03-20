package main

import "fmt"

func sayHello(firstName string, lastName string) {
	fmt.Println("Hello World", firstName, lastName)
}

func getHello(name string) string {
	return name
}

func getFullName(firstName string, lastName string) (string, string) {
	return firstName, lastName
}

func main() {
	result := getHello("Rafelino")
	sayHello(result, "Kelen")

	resultFirstName, resultLastName := getFullName("Rafelino", "Claudius Kelen")
	onlyFirstName, _ := getFullName("Rafelino", "")
	sayHello(resultFirstName, resultLastName)
	sayHello(onlyFirstName, "ck")
}
