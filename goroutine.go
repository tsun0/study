package main

import (
	"fmt"
	"time"
)

func main() {

	fmt.Println("main: start")

	fmt.Println("testを通常の関数として起動")
	test()

	fmt.Println("testをgoroutineとして起動")
	go test()

	time.Sleep(3 * time.Second)
	fmt.Println("main: finish")
}

func test() {

	for i := 0; i < 5; i++ {
		fmt.Println(i)
		time.Sleep(1 * time.Second)
	}
}
