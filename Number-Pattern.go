package main

import "fmt"

func main() {
	count:=0;
		for i:=1;i<=5;i++{
		    for j:=1;j<=i;j++{
		        count=count+1;
		        fmt.Print(count," ")
		    }
		    fmt.Println("")
		}
}
