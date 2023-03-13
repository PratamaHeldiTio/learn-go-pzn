package go_json

import (
	"encoding/json"
	"fmt"
	"os"
	"testing"
)

func TestJsonStreamDecoder(t *testing.T) {
	reader, _ := os.Open("customer.json")
	decoder := json.NewDecoder(reader) // parameternya harus type io.Reader

	customer := Customer{}
	decoder.Decode(&customer) // decode ini sama seperti unmarshall

	fmt.Println(customer)
}

func TestJsonStreamEncoder(t *testing.T) {
	writer, _ := os.Create("customer.json")
	encoder := json.NewEncoder(writer) // parameter harus type io.writer

	customer := Customer{
		Fisrtname: "Heldi",
		Lasname:   "botak",
		Age:       22,
	}

	_ = encoder.Encode(customer) // ini seperti json.Marshall

}
