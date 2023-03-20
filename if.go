package main

import "fmt"

func main() {
	var name string
	name = "doni"

	if name == "rafel" {
		fmt.Println("hello", name)
	} else {
		fmt.Println("sorry your not", name)
	}
}
