package main

import "fmt"

func main() {
	var names [3]string

	names[0] = "rafel"
	names[1] = "doni"
	names[2] = "teki"

	fmt.Println(names[0])

	var ages = [3]int{
		26, 27, 26,
	}
	fmt.Println(ages)
}
