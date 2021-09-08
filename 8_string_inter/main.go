package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

//global vars
var reader *bufio.Reader

type User struct {
	UserName  string
	Age       int
	FavNumber float64
}

func main() {
	reader = bufio.NewReader(os.Stdin)
	var user User
	user.UserName = readString("what is your name ?")
	user.Age = readInt("How old are you ?")
	fmt.Println("Your name is", user.UserName, " and you are ", user.Age, " years old!")
}

func prompt() {
	fmt.Print(" -> ")
}

func readString(s string) string {
	fmt.Println(s)
	prompt()
	userInput, _ := reader.ReadString('\n')
	userInput = strings.Replace(userInput, "\n", "", -1)
	return userInput
}

func readInt(s string) int {
	for {
		fmt.Println(s)
		prompt()
		userInput, _ := reader.ReadString('\n')
		userInput = strings.Replace(userInput, "\n", "", -1)
		num, err := strconv.Atoi(userInput)
		num = int(num)
		if err != nil {
			fmt.Println("Please enter a whole number!")
		} else {
			return num
		}
	}
}
