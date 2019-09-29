package main

import "fmt"

func main() {
	count := 10
	a:=0
	b := 1
	var sum int
	fmt.Print(a," ",b," ")
	for i := 2; i <= count; i++ {
		sum = a + b
		fmt.Print(sum," ")
		a = b
		b = sum
	}
	fmt.Println("")
}
