package main

import "fmt"

type Person2 struct {
	name string
	sex  byte
	age  int
}

func (tmp Person2) PrintInfo() {
	fmt.Printf("name=%s,sex=%c,age=%d\n", tmp.name, tmp.sex, tmp.age)
}

type Student2 struct {
	Person2
	id   int
	addr string
}

func (tmp Student2) PrintInfo() {
	fmt.Println("Student= ", tmp)
}

func main() {
	s := Student2{Person2{"liu", 'm', 12}, 12345, "beijing"}
	s.PrintInfo()
	s.Person2.PrintInfo()
}
