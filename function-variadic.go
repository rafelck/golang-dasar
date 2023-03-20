package main

import "fmt"

func sumAll(numbers ...int) int {
	total := 0
	for _, number := range numbers {
		total = total + number
	}

	return total
}

func main() {
	data := sumAll(1, 2)

	fmt.Println(data)

	numbers := []int{10, 10}
	total := sumAll(numbers...)
	fmt.Println(total)
}
