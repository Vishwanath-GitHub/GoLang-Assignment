package main
import "fmt"
func main() {
//defer uses LIFO Algorithm (Last In First Out).
//That means it works in stack concept.
//If we apply defer to all of the statements inside the function, everything will execute in reverse order
//So the last statement that is deferred will be the first one to be called
	defer fmt.Println("start")
	defer fmt.Println("middle")
	defer fmt.Println("end")
}
