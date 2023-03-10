package main

import (
	"fmt"
	"time"
)

func main() {
	now := time.Date(2023, 02, 16, 23, 59, 59, 0, time.Local)
	unixTime := now.Unix()

	fmt.Println(time.Unix(unixTime, 0))
}
