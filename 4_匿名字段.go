package main

import "fmt"

type Person struct {
	name string
	sex  byte
	age  int
}
type Student struct {
	Person
	id   int
	addr string
	name string
}

func main() {
	s1 := Student{Person{"mike", 'm', 18}, 1, "bj"}
	fmt.Printf("s2=%+v\n", s1)
}
