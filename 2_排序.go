package main

import "fmt"
import "math/rand"
import "time"

func mysort(a *[10]int) (result *[10]int) {
	len := len(*a)
	for i := 0; i < len; i++ {
		for j := 0; j < len-i-1; j++ {
			if (*a)[j] < (*a)[j+1] {
				(*a)[j], (*a)[j+1] = (*a)[j+1], (*a)[j]
			}
		}

	}
	return a
}

func main() {
	rand.Seed(time.Now().UnixNano())
	var a [10]int
	for i := 0; i < 10; i++ {
		a[i] = rand.Intn(100)
	}
	fmt.Println("a=", a)
	mysort(&a)
	fmt.Println("a=", a)
}
