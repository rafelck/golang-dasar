package main

import "fmt"

type Address struct {
	City, Province, Country string
}

func ChangeCountryToIndonesia(address *Address) {
	address.Country = "Indonesia"
}

func main() {
	var address1 Address = Address{
		City:     "subang",
		Province: "jawa barat",
		Country:  "Indonesia",
	}

	var address2 *Address = &address1

	address2 = &Address{
		City:     "Malang",
		Province: "jawa timur",
		Country:  "Indonesia",
	}

	var address3 = Address{
		City:     "Samarinda",
		Province: "Kalimantan timur",
		Country:  "",
	}

	ChangeCountryToIndonesia(&address3)

	fmt.Println(address1)
	fmt.Println(address2)
	fmt.Println(address3)

}
