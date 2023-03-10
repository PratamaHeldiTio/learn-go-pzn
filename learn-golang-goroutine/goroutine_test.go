package hello

import (
	"fmt"
	"testing"
)

func Hello() {
	fmt.Println("hello")
}

func TestHello(t *testing.T) {
	go Hello()
	fmt.Println("ups")
}
