package main

import "fmt"

func main() {

	// for variable bisa diakses diluar for
	counter := 1
	for counter <= 10 {
		fmt.Println("ke", counter)
		counter++
	}

	// for variabel hanya bisa diakses didalam for
	for counter := 1; counter <= 15; counter++ {
		fmt.Println("ke", counter)
	}
	fmt.Println(counter)

	names := []string{"baby", "heldi", "tio"}

	// for range with index
	for i, name := range names {
		fmt.Println("index", i, "nama", name)
	}

	// for range without index
	for _, name := range names {
		fmt.Println(name)
	}

	person := make(map[string]string)
	person["name"] = "heldi"

	fmt.Println(person["name"])

	fmt.Println(panggilan("mas"))

	name, umur := nameUmurNamed()
	fmt.Println(name, umur)
}

func panggilan(pang string) string {
	panggilan := pang + " heldi"
	return panggilan
}

// func multiple retur value

func nameUmur() (string, uint8) {
	return "heldi", 20
}

// menamakan return valuu
func nameUmurNamed() (name, asal string) {
	name = "tio"
	asal = "bdg"
	return
}
