package main

import (
	"time"
)

func main() {

	for i := 0; i <= 100; i++ {
		println("Hell o World  !!! Golang Start!!;", i)
		time.Sleep(50 * time.Millisecond)
	}
}
