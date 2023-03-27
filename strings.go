package main

import (
	"fmt"
	"strings"
)

func main() {
	fmt.Println(strings.Contains("Rafelino Claudius Kelen", "Clau"))
	fmt.Println(strings.Contains("Rafelino Claudius Kelen", "don"))

	fmt.Println(strings.Split("Rafelino Claudius Kelen", " "))

	fmt.Println(strings.ToLower("Rafelino Claudius Kelen"))
	fmt.Println(strings.ToUpper("Rafelino Claudius Kelen"))

	fmt.Println(strings.ToTitle("Rafelino Claudius Kelen"))

	fmt.Println(strings.ReplaceAll("Rafelino Claudius Kelen", "Rafelino", "Mr."))
}
