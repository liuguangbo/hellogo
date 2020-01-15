package main

import (
	"fmt"
	"runtime"
)

func test() {
	defer fmt.Println("cccccc")
	runtime.Goexit()
	fmt.Println("ddd")
}

func main() {

	go func() {
		fmt.Println("aaaaaaa")
		test()
		fmt.Println("bbbbbb")
	}()

	for {

	}
}
