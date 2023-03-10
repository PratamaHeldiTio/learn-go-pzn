package main

import (
	"errors"
	"fmt"
)

func Pembagian(nilai int, pembagi int) (int, error) {
	if pembagi == 0 {
		return 0, errors.New("gaboleh nol")
	} else {
		result := nilai / pembagi
		return result, nil
	}
}

func main() {
	hasil, err := Pembagian(10, 3)

	if err != nil {
		fmt.Println("err ", err.Error())
	} else {
		fmt.Println("hasil", hasil)
	}
	result := random()

	fmt.Println(result.(string))

}

func random() interface{} {
	return "dfdf"
}
