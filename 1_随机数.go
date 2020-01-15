package main

import "fmt"
import "math/rand"
import "time"

func main() {
	rand.Seed(time.Now().UnixNano())
	for i := 0; i < 5; i++ {
		fmt.Println("rand=", rand.Intn(30))
	}
}
