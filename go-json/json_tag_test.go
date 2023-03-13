package go_json

import (
	"encoding/json"
	"fmt"
	"testing"
)

type Notebook struct {
	Id       string `json:"id"`
	Brand    string `json:"brand"`
	ImageUrl string `json:"image_url"`
}

func TestJsonTag(t *testing.T) {
	notebook := Notebook{
		Id:       "P2424",
		Brand:    "Apple",
		ImageUrl: "https://sdfgfngof.com",
	}

	bytes, _ := json.Marshal(notebook)

	fmt.Println(string(bytes))
}
