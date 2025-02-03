package main

import "fmt"

type Person struct {
	id int
	name string
	age int
}

func (this *Person) Eat() {
	fmt.Println("person is eating")
}

func (this *Person) Sleep() {
	fmt.Println("person is sleeping")
}

type Employee struct {
	Person
	salary int
}

func (this *Employee) Sleep() {
	fmt.Println("employee is sleeping")
}

func main() {
	p := Person{1, "Alice", 25}
	p.Eat()
	p.Sleep()

	e := Employee{Person{2, "Bob", 30}, 5000}
	// !! err -> e := Employee{Person2, "Bob", 30, 5000}
	e.Eat()
	e.Sleep()
}