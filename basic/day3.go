package basic

import "fmt"

type Person struct {
	name    string
	age     int
	address string
}

type student struct {
	Person
	school string
	class  string
}

func (p *Person) greet() {
	fmt.Println("Hello, ", p.name, p.age)
}

func (p *Person) setAddress(newAddress string) {
	p.address = newAddress
}

func (s *student) setClass(class string) {
	s.class = class
}

type Animal interface {
	run() string
}

type Dog struct {
	name string
}

type Cat struct {
	name string
}

func (d *Dog) run() string {
	return d.name + "is running"
}

func (c *Cat) run() string {
	return c.name + "is running"
}
