// package declaration
package main

//library/package imports
import "fmt"

func main() {
	//format
	fmt.Println("hello_world")
	sayHelloWorld("byebye")
}

// parameter_name type
func sayHelloWorld(whatToSay string) {
	fmt.Println(whatToSay)
}

/*
RUN PROGRAM
go run .\main.go
*/
