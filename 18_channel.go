package main

import (
	"fmt"
	"time"
)

var ch = make(chan int)

func Printer(str string) {
	for _, data := range str {
		fmt.Printf("%c", data)
		time.Sleep(time.Second)
	}
	fmt.Println("\n")
}
func person1() {
	Printer("hello")
	ch <- 666
}
func person2() {
	<-ch //从管道取数据，如果通道没有数据将阻塞
	Printer("world")
}

func main() {
	go person1()
	go person2()
	for {

	}
}
