package main

import "fmt"

func main() {
	const name = "heldi"
	fmt.Println(name)
	// name = "tio" // error

	// multiple const
	const (
		firstName  string = "Heldi"
		middleName        = "tio"
	)

	fmt.Println(firstName)
}
