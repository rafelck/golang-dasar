package main

import (
	"BELAJAR-GOLANG-DASAR/database"
	"fmt"
)

func main() {
	var result string
	result = database.GetDatabase()
	fmt.Println(result)
}
