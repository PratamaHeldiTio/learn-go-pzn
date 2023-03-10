package main

import (
	"context"
	"fmt"
	"runtime"
	"sync"
	"testing"
	"time"
)

func TestContextWithValue(t *testing.T) {
	contextA := context.Background()

	contextB := context.WithValue(contextA, "b", "B")
	contextC := context.WithValue(contextA, "c", "C")

	contextD := context.WithValue(contextB, "d", "D")
	contextE := context.WithValue(contextB, "e", "E")

	contextF := context.WithValue(contextC, "f", "F")
	contextG := context.WithValue(contextC, "g", "G")

	fmt.Println(contextA)
	fmt.Println(contextB)
	fmt.Println(contextC)
	fmt.Println(contextD)
	fmt.Println(contextE)
	fmt.Println(contextF)
	fmt.Println(contextG)

	fmt.Println(contextF.Value("f"))
	fmt.Println(contextF.Value("c"))
	fmt.Println(contextF.Value("b"))
	fmt.Println(contextA.Value("b"))
}

func CreateConter(ctx context.Context, group *sync.WaitGroup) chan int {
	destination := make(chan int)

	go func() {
		defer close(destination)
		defer group.Done()
		group.Add(1)

		counter := 0

		for {
			select {
			case <-ctx.Done(): // ketika mendapatkan sinyal cancel
				return
			default:
				counter += 1
				destination <- counter
				time.Sleep(1 * time.Second) // simulasi slow
			}
		}
	}()

	return destination
}

func TestCtxCancel(t *testing.T) {
	fmt.Println("awal", runtime.NumGoroutine())

	group := &sync.WaitGroup{}
	ctxParent := context.Background()            // context parent
	ctx, cancel := context.WithCancel(ctxParent) // child context withcancel

	destination := CreateConter(ctx, group)

	for n := range destination {
		fmt.Println("counter", n)

		if n == 10 {
			break
		}
	}

	cancel() // mengirim sinyal cancel

	group.Wait()
	fmt.Println("awal", runtime.NumGoroutine())
}

func TestCtxTimeout(t *testing.T) {
	fmt.Println("awal", runtime.NumGoroutine())

	group := &sync.WaitGroup{}
	ctxParent := context.Background()
	ctx, cancel := context.WithTimeout(ctxParent, 5*time.Second)
	defer cancel() // ketida proses lebih cepat dari timeout go routine di cancel

	destination := CreateConter(ctx, group)

	for n := range destination {
		fmt.Println("counter", n)
	}

	group.Wait()
	fmt.Println("awal", runtime.NumGoroutine())
}
