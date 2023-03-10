package hello

import (
	"fmt"
	"strconv"
	"testing"
	"time"
)

func TestChannel(t *testing.T) {
	channelName := make(chan string)
	channelAge := make(chan uint8)
	defer close(channelName)
	defer close(channelAge)

	fmt.Println("ini  awal bgt")

	go GivemeResponseHeldi(channelName)
	fmt.Println("ini  antara panggil goroutine")

	go GivemeResponseUmur(channelAge)

	fmt.Println("ini diantara channel")

	name := <-channelName
	age := <-channelAge
	fmt.Println("ini diantara channel")

	fmt.Println("ini channelnya", name)

	fmt.Println("ini channelnya", name, age)

	fmt.Println("ini dibawah print channel")

	time.Sleep(2 * time.Second)
}

func GivemeResponseHeldi(channel chan<- string) {
	fmt.Println("sebelum kirim channel nama")
	channel <- "heldi"
	fmt.Println("setelah kirim channel nama")
}

func GivemeResponseUmur(channel chan uint8) {
	fmt.Println("sebelum kirim channel umur")
	channel <- 20
	fmt.Println("setelah kirim channel umur")
}

func TestBufferedChhannel(t *testing.T) {
	channel := make(chan string, 2)
	defer close(channel)

	channel <- "heldi"
	channel <- "tio"

	fmt.Println(<-channel)
	fmt.Println(<-channel)

}

func TestRange(t *testing.T) {
	channel := make(chan string)

	go func() {
		for i := 0; i < 10; i++ {
			fmt.Println("dalam goroutine")
			channel <- "perulangan ke " + strconv.Itoa(i)
		}
		close(channel)
	}()

	for data := range channel {
		fmt.Println("dalam for")
		fmt.Println("menerima data", data)
	}
}

func TestSelect(t *testing.T) {
	channel1 := make(chan string)
	channel2 := make(chan string)
	defer close(channel1)
	defer close(channel2)

	go GivemeResponseHeldi(channel1)
	go GivemeResponseHeldi(channel2)

	counter := 0
	for {
		select {
		case data := <-channel1:
			fmt.Println("data channel 1", data)
			counter++
		case data := <-channel2:
			fmt.Println("data channel 2", data)
			counter++
		default:
			fmt.Println("menunggu data")
		}
		if counter == 2 {
			break
		}
	}

}
