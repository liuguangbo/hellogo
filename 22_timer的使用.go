package main

import (
	"fmt"
	"time"
)

func main() {
	timer := time.NewTimer(5 * time.Second)
	fmt.Println("当前时间: ", time.Now())
	t := <-timer.C
	fmt.Println("t= ", t)
}
