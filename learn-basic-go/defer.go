package main

import "fmt"

func main() {
	runApplication(false)
	fmt.Println("check")
}

func runApplication(err bool) {
	defer logging()
	if err {
		panic("err")
	}
	fmt.Println("run App")
}

func logging() {
	message := recover()
	fmt.Println(message)
	fmt.Println("logging")
}
