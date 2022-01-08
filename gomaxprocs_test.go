package goroutines

import (
	"fmt"
	"runtime"
	"testing"
)

func TestGet(t *testing.T) {
	totalCPU := runtime.NumCPU()
	fmt.Println(totalCPU)

	totalThread := runtime.GOMAXPROCS(-1)
	fmt.Println(totalThread)

	totalGoRoutines := runtime.NumGoroutine()
	fmt.Println(totalGoRoutines)
}
