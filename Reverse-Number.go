package main

import (
	"fmt"
)

func reverse_number(no int) int {
	result := 0
	for no > 0 {
		rem := no % 10
		result = result * 10
		result = result + rem
		no = no / 10
	}
	return result
}

func main() {
	fmt.Println(reverse_number(478315719))
}
