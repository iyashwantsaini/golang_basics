package main

import (
	"bufio"
	"os"
	"strings"

	//appName in go.mod / package name
	"doctorapp/doctor"
	"fmt"
)

func main() {
	//user input reader object
	reader := bufio.NewReader(os.Stdin)

	var whttosay string
	whttosay = doctor.Intro()
	fmt.Println(whttosay)

	//for loop
	for i := 0; i < 5; i++ {
		userInput, _ := reader.ReadString('\n')
		//for windows users strip enter press from input
		userInput = strings.Replace(userInput, "\r\n", "", -1)
		//for other OS -> linux, mac
		userInput = strings.Replace(userInput, "\n", "", -1)
		if userInput == "quit" {
			break
		} else {
			// fmt.Println(userInput)
			response := doctor.Response(userInput)
			fmt.Println(response)
		}
	}
}

/*
BUILD A GO FILE TO EXE
go build -o doctor.exe main.go
*/
