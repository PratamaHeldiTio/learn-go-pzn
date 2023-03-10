package hello

import (
	"fmt"
	"sync"
	"testing"
	"time"
)

func AddToMap(value int, data *sync.Map, group *sync.WaitGroup) {
	defer group.Done()
	group.Add(1)
	data.Store(value, value)
}

func TestMap(t *testing.T) {
	data := &sync.Map{}
	group := &sync.WaitGroup{}

	for i := 0; i < 100; i++ {
		go AddToMap(i, data, group)
	}

	group.Wait()

	data.Range(func(key, value interface{}) bool {
		fmt.Println(key, ":", value)
		return true
	})

	fmt.Println("Done")
}

// cond merupakan skema looking akan tetapi bisa menunggu melakukan perintah yang dilock asalkan mendapatkan signal

var loocker = sync.Mutex{}        // membuat loocker
var cond = sync.NewCond(&loocker) // membuat cond
var group = sync.WaitGroup{}      // membuat wait group agar menunggu go routine selesai

func WaitCondition(value int) {
	defer group.Done()
	group.Add(1)

	cond.L.Lock() // melakukan locking
	cond.Wait()   // menunggu perintah sinyal untuk melakukan
	fmt.Println("cond", value)
	cond.L.Unlock() // unlock
}

func TestCond(t *testing.T) {
	for i := 0; i < 10; i++ {
		go WaitCondition(i)
	}

	for i := 0; i < 10; i++ {
		time.Sleep(1 * time.Second)
		cond.Signal()
	}

	group.Wait()
}
