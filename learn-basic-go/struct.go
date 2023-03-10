package main

import "fmt"

type Customer struct {
	Name string
	Age  uint8
}

func (c Customer) SayHello(name string) {
	fmt.Println("helo " + name + " my name is " + c.Name)
}

func (c Customer) SayHi(name string) {
	fmt.Println("hai " + name + " my name is " + c.Name)
}

func main() {
	eko := Customer{
		Name: "eko",
		Age:  30,
	}

	heldi := Customer{
		Name: "heldi",
		Age:  30,
	}

	eko.SayHello("heldi")
	heldi.SayHello("joko")

	var aku interface{} = nil
	fmt.Println(aku)

}
