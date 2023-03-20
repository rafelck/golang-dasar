package main

import "fmt"

func main() {

	for counter := 1; counter <= 10; counter++ {
		fmt.Println(counter)
	}

	slice := []string{"rafel", "idang", "ribka"}

	for i := 0; i < len(slice); i++ {
		fmt.Println(slice[i])
	}

	for i, name := range slice {
		fmt.Println("Index", i, "=", name)
	}
}
