package main

import (
	"fmt"
	"runtime"
	"time"
)

func main() {
	for i := 0; i <= 100; i++ {
		println("Hell o World  !!! Go ", i)
		time.Sleep(1 * time.Millisecond)
	}

	fmt.Println("cpu count = ", runtime.NumCPU())

}
