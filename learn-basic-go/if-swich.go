package main

import "fmt"

func main() {
	if name := "heldi"; name == "tio" {
		fmt.Println("bukan tio")
	} else if name == "heldi" {
		fmt.Println("iyaa heldi")
		if umur := 20; umur == 20 {
			fmt.Println("uurnya", umur)
		}
	} else {
		fmt.Println("gakenal")
		// fmt.Println(umur)
	}

	switch name := "heldi"; name {
	case "tio":
		fmt.Println("bukan tio")
	case "heldi":
		fmt.Println("iyaa heldi")
		switch umur := 20; umur {
		case 20:
			fmt.Println("uurnya", umur)
		}
	default:
		fmt.Println("gakenal")
	}

	name := "heldi"
	// switch no variable
	switch {
	case name == "tio":
		fmt.Println("bukan tio")
	case name == "heldi":
		fmt.Println("iyaa heldi")
	default:
		fmt.Println("gakenal")
	}
}
