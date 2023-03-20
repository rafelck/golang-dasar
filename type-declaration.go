package main

import "fmt"

func main() {
	type noKtp string
	type married bool

	var noKtpRafel noKtp = "10320340230"
	var statusMenikah married = true
	fmt.Println(noKtpRafel)
	fmt.Println(statusMenikah)
}
