package main

import (
	"fmt"
	"math/rand"
	"time"
)

func Initdata(s []int) {
	rand.Seed(time.Now().UnixNano())
	a := len(s)
	for i := 0; i < a; i++ {
		s[i] = rand.Intn(100)
	}
}

func Mysort(a []int) {
	len := len(a)
	for i := 0; i < len; i++ {
		for j := 0; j < len-i-1; j++ {
			if (a)[j] < (a)[j+1] {
				(a)[j], (a)[j+1] = (a)[j+1], (a)[j]
			}
		}

	}
	fmt.Println("a=", a)
}

func main() {
	s := make([]int, 10)
	Initdata(s)
	fmt.Println("s=", s)
	fmt.Printf("\n")
	Mysort(s)

	fmt.Println("s=", s)

}
