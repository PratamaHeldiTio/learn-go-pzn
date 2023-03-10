package main

import "fmt"

func main() {
	var (
		nilai32 int32 = 200
		nilai16 int16 = int16(nilai32)
		nilai8  int8  = int8(nilai32)
	)

	fmt.Println(nilai32)
	fmt.Println(nilai16)
	fmt.Println(nilai8)

	// konversi byte karakter kestring
	var name = "heldi"
	e := name[0]
	estring := string(e)
	fmt.Println(estring)
}
