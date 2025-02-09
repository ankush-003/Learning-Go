package main

import (
	"fmt"
)

type Person struct {
	Name string
	Health int
}

func (p Person) TakeDamage(amount int) {
	p.Health -= amount
}

func (p *Person) TakeRealDamage(amount int) {
	p.Health -= amount
}

func TakeRealDamage(p *Person, amount int) { // same as above
	p.Health -= amount
}

func main() {
	var p Person
	p.Name = "Ankush"
	p.Health = 100
	fmt.Println(p.Health)
	p.TakeDamage(10)
	 
	fmt.Println(p.Health)
	p.TakeRealDamage(10)
	fmt.Println(p.Health)
	TakeRealDamage(&p, 10)
	fmt.Println(p.Health)
}