package main

import (
	"fmt"
)

func main() {
	showType(nil)
	showType(123)

	showType2(nil)
	showType2(123)
}

func showType(x interface{}) {
	switch x.(type) {
	case nil:
		fmt.Println("type: nil")
	case int:
		fmt.Println("type: int")
	default:
		fmt.Println("type: undifined")
	}
	fmt.Println("aaaa")
}

func showType2(x interface{}) {
	switch x.(type) {
	case nil:
		fmt.Println("type: nil")
		return
	case int:
		fmt.Println("type: int")
		return
	default:
		fmt.Println("type: undifined")
		return
	}
	fmt.Println("aaaa")
}
