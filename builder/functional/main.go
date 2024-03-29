package main

import "fmt"

type Person struct {
	name, position string
}

type (
	personMod     func(*Person)
	PersonBuilder struct {
		actions []personMod
	}
)

func (b *PersonBuilder) Called(name string) *PersonBuilder {
	b.actions = append(b.actions, func(p *Person) {
		p.name = name
	})

	return b
}

func (b *PersonBuilder) WorksAsA(position string) *PersonBuilder {
	b.actions = append(b.actions, func(p *Person) {
		p.position = position
	})

	return b
}

func (b *PersonBuilder) Build() *Person {
	p := Person{}
	for _, a := range b.actions {
		a(&p)
	}

	return &p
}

func main() {
	b := PersonBuilder{}
	p := b.Called("Angelo").WorksAsA("Engineer").Build()

	fmt.Println(*p)
}
