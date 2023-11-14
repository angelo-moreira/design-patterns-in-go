package main

import "github.com/davecgh/go-spew/spew"

type Employee struct {
	Name, Position string
	AnnualIncome   int
}

// functional
func NewEmployeeFactory(position string, annualIncome int) func(name string) *Employee {
	return func(name string) *Employee {
		return &Employee{name, position, annualIncome}
	}
}

// structural

type EmployeeFactory struct {
	Position     string
	AnnualIncome int
}

func (f *EmployeeFactory) Create(name string) *Employee {
	return &Employee{name, f.Position, f.AnnualIncome}
}

func NewEmployeeFactory2(position string, annualIncome int) *EmployeeFactory {
	return &EmployeeFactory{position, annualIncome}
}

func main() {
	developerFactory := NewEmployeeFactory("developer", 60000)
	managerFactory := NewEmployeeFactory("manager", 80000)

	developer := developerFactory("Adam")
	manager := managerFactory("Jane")

	spew.Dump(developer)
	spew.Dump(manager)

	bossFactory := NewEmployeeFactory2("CEO", 100000)
	boss := bossFactory.Create("Sam")
	spew.Dump(boss)
}
