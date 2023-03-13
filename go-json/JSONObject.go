package go_json

import (
	"encoding/json"
	"fmt"
	"testing"
)

type Address struct {
	City   string
	Street string
}

type Customer struct {
	Fisrtname string
	Lasname   string
	Age       int
	Married   bool
	Hobbies   []string
	Address   []Address
}

func TestJSONObject(t *testing.T) {
	customer := Customer{
		Fisrtname: "Heldi",
		Lasname:   "Pratama",
		Age:       22,
		Married:   false,
	}

	bytes, err := json.Marshal(customer)
	if err != nil {
		panic(err)
	}

	fmt.Println(string(bytes))
}
