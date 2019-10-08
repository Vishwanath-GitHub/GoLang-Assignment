package main

import (
	"fmt"
	"os"
)

func main() {
	usingdefer1()
	fmt.Println()
	usingdefer2()
	fmt.Println()
	// Immediately after getting a file object with
	// `createFile`, we defer the closing of that file
	// with `closeFile`. This will be executed at the end
	// of the enclosing function (`main`), after
	// `writeFile` has finished.
	// Suppose we wanted to create a file, write to it,
	// and then close when we're done. Here's how we could
	// do that with `defer`.
	f := createFile("C:/test/hello.txt")
	defer closeFile(f)
	writeFile(f)
}
/*defer 'keyword' executes any function defined or declared with it after the current function finishes its final
statement but before it actually returns.*/
/*deferring doesn't move anything defined to the end of the function,
 it actually moves to after the current function before it returns*/
//It can be regarded as 'finally' keyword of GoLang
//defer can be used anywhere in a function (including main function since it is also a function)
//In the defer statements, the arguments are evaluated when the defer statement executed, not when they called.
func usingdefer1 () {
	fmt.Println("start")
	defer fmt.Println("middle") //notice in the result that this line executes at the end of the function
	fmt.Println("end")
}
//defer uses LIFO Algorithm (Last In First Out).
//That means it works in stack concept.
//If we apply defer to all of the statements inside the function, everything will execute in reverse order
//So the last statement that is deferred will be the first one to be called
func usingdefer2 () {
	defer fmt.Println("start")
	defer fmt.Println("middle")
	defer fmt.Println("end")
}
func createFile(p string) *os.File {
	fmt.Println("creating")
	f, err := os.Create(p)
	if err != nil {
		panic(err)
	}
	return f
}

func writeFile(f *os.File) {
	fmt.Println("writing")
	fmt.Fprintln(f, "data") //this is writing 'data' keyword in the file by replacing everything inside it

}

func closeFile(f *os.File) {
	fmt.Println("closing")
	err := f.Close()
	// It's important to check for errors when closing a
	// file, even in a deferred function.
	if err != nil {
		fmt.Fprintf(os.Stderr, "error: %v\n", err)
		os.Exit(1)
	}
}
