package main

import (
	"fmt"
	"io/ioutil"
	"math/rand"
	"runtime"
	"time"
)

func main() {
	var b []byte
	var err error

	rand.Seed(time.Now().UnixNano())
	rand := rand.Intn(100)

	for i := 0; i <= 10; i++ {
		println("Hell o World  !!! Go ", i)
		time.Sleep(1 * time.Millisecond)
	}

	fmt.Println("cpu count = ", runtime.NumCPU())

	b, err = ioutil.ReadFile("./test.txt")

	if err == nil {
		fmt.Printf("%s", b)
	}

	if b, err := ioutil.ReadFile("./test.txt"); err == nil {
		fmt.Printf("%s", b)
	}

	fmt.Println(rand)

	for i := 0; i < 100; i++ {

		switch {
		case i%3 == 0 && i%5 == 0:
			fmt.Printf("%d = 3과 5의 공배수\r\n", i)
		case i%3 == 0:
			fmt.Printf("%d = 3의배수\r\n", i)
		case i%5 == 0:
			fmt.Printf("%d =5의 배수\r\n", i)

		}
	}

}
