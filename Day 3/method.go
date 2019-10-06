package main

import "fmt"

func main() {

	g:=greeter{
		greeting: "Hello",
		name: "Go",
	}
	g.greet()
	g.greet1()
	fmt.Println("The new name is: ",g.name)
}
type greeter struct {
	greeting string
	name string
}

func (g greeter) greet()  {
	fmt.Println(g.greeting,g.name)
}

func (g *greeter) greet1() {
	fmt.Println(g.greeting,g.name)
	g.name="Vishu"
}
