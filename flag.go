package main

import (
	"flag"
	"fmt"
)

func main() {
	var host *string = flag.String("host", "localhost", "Put your database host")
	var user *string = flag.String("user", "root", "Put your database user")
	var password *string = flag.String("password", "root", "Put your database password")

	flag.Parse()

	fmt.Println(*host)
	fmt.Println(*user)
	fmt.Println(*password)
}
