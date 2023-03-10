package hello

import (
	"fmt"
	"sync"
	"testing"
)

func TestPoolTest(t *testing.T) {
	pool := sync.Pool{
		New: func() interface{} {
			return "tunggu"
		},
	}

	group := &sync.WaitGroup{}

	pool.Put("heldi")
	pool.Put("tio")
	pool.Put("pratama")

	for i := 0; i < 20; i++ {
		go func() {
			defer group.Done()
			group.Add(1)
			data := pool.Get()
			fmt.Println(data)
			pool.Put(data)
		}()
	}

	group.Wait()
	fmt.Println("selesai")
}
