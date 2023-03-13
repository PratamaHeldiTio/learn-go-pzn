package go_json

import (
	"encoding/json"
	"fmt"
	"testing"
)

func TestDecode(t *testing.T) {
	stringJSON := `{"Fisrtname":"Heldi","Lasname":"Pratama","Age":22,"Married":false}`
	byteJSON := []byte(stringJSON)
	customer := Customer{}

	if err := json.Unmarshal(byteJSON, &customer); err != nil {
		panic(err)
	}

	fmt.Println(customer)
	fmt.Println(customer.Fisrtname)
	fmt.Println(customer.Lasname)
	fmt.Println(customer.Age)
	fmt.Println(customer.Married)
}
