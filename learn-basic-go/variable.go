package main

import "fmt"

func main() {
	var name string

	name = "heldi"
	fmt.Println(name)

	name = "tio"
	fmt.Println(name)

	// name = 100  // error tipe data beda

	var friendName = "hendra"
	fmt.Println(friendName)

	var age uint8 = 15
	fmt.Println(age)

	// tidak wajib var

	address := "tanjungsari"
	fmt.Println(address)

	phone := 15
	fmt.Println(phone)

	var (
		ibu  = "ibu"
		ayah = "ayah"
	)
	fmt.Println(ibu, ayah)
}
