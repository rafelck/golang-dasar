package main

import "fmt"

func main() {
	var person = map[string]string{
		"name": "Rafelino",
		"age":  "26",
	}

	person["title"] = "programmer"

	fmt.Println(person)
	fmt.Println(person["name"])
	fmt.Println(person["age"])

	var book map[string]string = make(map[string]string)

	book["title"] = "belajar golang"
	book["author"] = "rafelino"
	book["ups"] = "salah"

	fmt.Println(book)
	fmt.Println(len(book))

	delete(book, "ups")

	fmt.Println(book)
	fmt.Println(len(book))
}
