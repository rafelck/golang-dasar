package main

import "fmt"

func main() {
	var months = [...]string{
		"januari",
		"febuari",
		"maret",
		"april",
		"mei",
		"juni",
		"july",
		"agustus",
		"september",
		"oktober",
		"november",
		"desember",
	}

	var slice1 = months[4:7]
	fmt.Println(slice1)
	fmt.Println(len(slice1))
	fmt.Println(cap(slice1))
}
