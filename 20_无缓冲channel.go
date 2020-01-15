package main

import (
	"fmt"
	"time"
)

func main() {
	ch := make(chan int)
	fmt.Printf("len(ch)=%d,cap(ch)=%d\n", len(ch), cap(ch))
	go func() {
		for i := 0; i < 3; i++ {
			fmt.Println("子协程:i= ", i)
			ch <- i
		}
	}()
	time.Sleep(5 * time.Second)
	for i := 0; i < 3; i++ {
		num := <-ch
		fmt.Println("num= ", num)
	}
}
