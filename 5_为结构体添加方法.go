package main

import (
	"fmt"
)

type student struct {
	id   int
	age  int
	name string
}

func (tmp student) PrintInfo() {
	fmt.Println("tmp= ", tmp)
}

func main() {
	a := student{2, 32, "liu"}
	a.PrintInfo()

}
