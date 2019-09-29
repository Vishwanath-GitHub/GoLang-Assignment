package main
import "fmt"
func main()  {
	for i:=1;i<=5;i++{
		    for j:=4;j>i-1;j--{
		        fmt.Print(" ");
		    }
		    for k:=1; k<=i; k++{
                fmt.Print("*");
            }
            fmt.Println("");
	}
}
