package go_json

import (
	"encoding/json"
	"fmt"
	"testing"
)

func TestJsonArrayEncode(t *testing.T) {
	customer := Customer{
		Fisrtname: "Heldi",
		Lasname:   "Pratama",
		Age:       22,
		Married:   false,
		Hobbies:   []string{"game", "reading", "making"},
	}

	bytes, _ := json.Marshal(customer)

	fmt.Println(string(bytes))
}

func TestJsonArrayDecode(t *testing.T) {
	jsonString := `{"Fisrtname":"Heldi","Lasname":"Pratama","Age":22,"Married":false,"Hobbies":["game","reading","making"]}`
	jsonByte := []byte(jsonString)
	customer := Customer{}

	if err := json.Unmarshal(jsonByte, &customer); err != nil {
		panic(err)
	}

	fmt.Println(customer)
	fmt.Println(customer.Fisrtname)
	fmt.Println(customer.Lasname)
	fmt.Println(customer.Age)
	fmt.Println(customer.Married)
	fmt.Println(customer.Hobbies)
}

func TestJsonArrayComplex(t *testing.T) {
	customer := Customer{
		Fisrtname: "Heldi",
		Lasname:   "Pratama",
		Age:       22,
		Married:   false,
		Hobbies:   []string{"game", "reading", "making"},
		Address: []Address{
			{
				City:   "sumedang",
				Street: "tjs",
			},
			{
				City:   "bangkalan",
				Street: "kamal",
			},
		},
	}

	bytes, _ := json.Marshal(customer)

	fmt.Println(string(bytes))

}
