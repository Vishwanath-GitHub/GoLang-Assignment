package main

import "fmt"

type person struct {
	n string
	age int
}

type intf interface {
	display()
	set()
	get()
}

func (p *person) get() {
	fmt.Println(p.n,p.age)
}

func (p *person) set() {
	fmt.Println(p.n,p.age)
}

func (p *person) display() {

}
func main() {
	fmt.Println("hello")
	per:=person{"vishu",21}
	per.get()
	per.set()
	per.display()
}
