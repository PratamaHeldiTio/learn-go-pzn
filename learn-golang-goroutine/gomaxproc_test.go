package hello

import (
	"fmt"
	"runtime"
	"testing"
)

func TestGoMax(t *testing.T) {
	totalCpu := runtime.NumCPU()
	fmt.Println(totalCpu)

	totalTh := runtime.GOMAXPROCS(-1)
	fmt.Println(totalTh)

	totalGo := runtime.NumGoroutine()
	fmt.Println(totalGo)
}
