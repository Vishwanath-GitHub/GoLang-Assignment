package main

import (
	"fmt"
	"net/http"
)

func main() {
	//GoLang doesn't have exceptions
	//Panic is a situation where our application can't continue to function and it can't figure out what to do
	/*When a function encounters a panic, its execution is stopped, any deferred functions are executed
	and then the control returns to its caller.
	This process continues until all the functions of the current goroutine have returned at which point the program
	prints the panic message,
	followed by the stack trace and then terminates.
	 */
	var1,var2:=5,0
	var3:=var1/var2
	fmt.Println(var3)
	//panic function is used so that we get something in response as we desired by passing it in the arugument
	fmt.Println("Just started")
	panic("It stopped")

	fmt.Println("It didn't stop")
	firstName := "Vishu"
	fullName(&firstName, nil)
	fmt.Println("returned normally from main")

	//This will create a http localhost service displaying 'Hello there'
	//Panic will arrive if it fails to do so
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("Hello there"))
	})
	err:=http.ListenAndServe(":8080",nil)
	if err != nil {
		panic(err.Error())
	}
}
func fullName(firstName *string, lastName *string) {
	if firstName == nil {
		panic("runtime error: first name cannot be nil")
	}
	if lastName == nil {
		panic("runtime error: last name cannot be nil")
	}
	fmt.Printf("%s %s\n", *firstName, *lastName)
	fmt.Println("returned normally from fullName")
}
