package hello

import (
	"fmt"
	"sync"
	"testing"
)

func RunAsync(group *sync.WaitGroup, loop int) {
	defer group.Done()
	group.Add(1)

	fmt.Println("hello", loop)

}

func TestWaitGroup(t *testing.T) {
	group := &sync.WaitGroup{}

	for i := 0; i < 100; i++ {
		go RunAsync(group, i)
	}

	group.Wait()
	fmt.Println("done")
}
