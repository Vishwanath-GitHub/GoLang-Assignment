package main

import "fmt"

func main() {
	//recover is like the 'catch' of try-catch concept
	//recover is a builtin function which is used to regain control of a panicking goroutine.
/* Steps for recover function to work:
	1. When any panic occurs from a statement or a function, all the statements after it do not get executed.
	2. Then this panic searches for recover function in the same function and executes it to recover from panic & resumes
	   program execution.
	3. If it doesn't find any recover function inside same function, it searches the function that is holding the
	   recover function in itself and recovers from panic. This function which contains the recover function needs to be
	   deferred in the current function.
	 */
	defer fmt.Println("deferred call in main")
	firstName := "Elon"
	fullName(&firstName, nil)
	fmt.Println("returned normally from main")
}
func recoverName() {
	if r := recover(); r!= nil {
		fmt.Println("recovered from ", r)
	}
}
func fullName(firstName *string, lastName *string) {
	//when the panic from the if conditions arrives, it moves recoverName function
	//the recoverName function helps in recovery from panic
	//so the program execution continues
	defer recoverName()
	if firstName == nil {
		panic("runtime error: first name cannot be nil")
	}
	if lastName == nil {
		panic("runtime error: last name cannot be nil")
	}
	fmt.Printf("%s %s\n", *firstName, *lastName)
	fmt.Println("returned normally from fullName")
}
