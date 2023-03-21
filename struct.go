package main

import "fmt"

type Customer struct {
	Name, Address string
	Age           int
}

func (customer Customer) sayHi(name string) {
	fmt.Println("Hello", name, "My Name Is", customer.Name)
}

func main() {
	var customer Customer

	customer.Name = "Rafel"
	customer.Address = "Indonesia"
	customer.Age = 26

	fmt.Println(customer.Name)
	fmt.Println(customer.Address)
	fmt.Println(customer.Age)

	joko := Customer{
		Name:    "Doni",
		Address: "Yogyakarta",
		Age:     26,
	}
	fmt.Println(joko)
	joko.sayHi("Tono")
}
