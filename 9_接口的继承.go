package main

import "fmt"

type Humaner interface {
	sayhi()
}

type Person interface {
	Humaner
	sing(lrc string)
}

type Student struct {
	name string
	id   int
}

func (tmp *Student) sayhi() {
	fmt.Printf("student[%s,%d] sayhi \n", tmp.name, tmp.id)
}

func (tmp *Student) sing(lrc string) {
	fmt.Printf("student在吟唱", lrc)
}

func main() {
	var i Person
	s := &Student{"mike", 666}
	i = s
	i.sayhi()
	i.sing("是的风景")
}
