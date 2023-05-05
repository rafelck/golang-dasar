package main

import "fmt"

func main() {
	var name string

	name = "rafel"

	switch name {
	case "eko":
		fmt.Println("halo eko")
	case "rafel":
		fmt.Println("halo rafel")
	case "budi":
		fmt.Println("halo budi")
	default:
		fmt.Println("Who are you?")
	}
}
