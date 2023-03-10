package main

import "fmt"

type Address struct {
	City, Prov string
}

func main() {
	address1 := Address{"smd", "jabar"}
	// address2 := &address1
	address3 := &address1
	address4 := &address1
	aad := Address{}

	// address2.City = "bdg"

	address2 := &Address{"sm", "jabar"}

	coba(&address1)
	fmt.Println(address1)
	fmt.Println(*address2)
	fmt.Println(*address3)
	fmt.Println(*address4)
	fmt.Println(aad)
}

func coba(add *Address) {
	add.City = "heldi"
}
