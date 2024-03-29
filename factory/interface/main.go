package main

import "fmt"

type Person interface {
	SayHello()
}

type person struct {
	name string
	age  int
}

type tiredPerson struct {
	name string
	age  int
}

func (p *person) SayHello() {
	fmt.Printf("Hi my name is %s, I am %d years old \n", p.name, p.age)
}

func (tp *tiredPerson) SayHello() {
	fmt.Printf("Hi my name is %s, I am %d years old and I AM VERY TIRED \n", tp.name, tp.age)
}

func NewPerson(name string, age int) Person {
	if age > 100 {
		return &tiredPerson{name, age}
	}

	return &person{name, age}
}

func main() {
	p := NewPerson("Angelo", 142)
	p.SayHello()
}
