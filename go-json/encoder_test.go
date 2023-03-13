package go_json

import (
	"encoding/json"
	"fmt"
	"testing"
)

func logJson(data interface{}) {
	bytes, err := json.Marshal(data)
	if err != nil {
		panic(err)
	}

	fmt.Println(string(bytes))
}

func TestEncoder(t *testing.T) {
	logJson("heldi")
	logJson(124)
	logJson(true)
	logJson([]string{"heldi", "elsa", "sifa"})
}
