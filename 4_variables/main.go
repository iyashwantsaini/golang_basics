package main

import (
	"bufio"
	"fmt"
	"math/rand"
	"os"
	"time"
)

const prompt = " and press ENTER when ready."

func getans(firstNum int, secondNum int, subtraction int) {
	reader := bufio.NewReader(os.Stdin)

	fmt.Println("Guessing Game!")
	fmt.Println("--------------")

	fmt.Println("Enter number between 1 and 10 ", prompt)
	reader.ReadString('\n') //i.e stop reading after pressing enter

	fmt.Println("Multiply your num by ", firstNum, prompt)
	reader.ReadString('\n')

	fmt.Println("Multiply result by ", secondNum, prompt)
	reader.ReadString('\n')

	fmt.Println("Divide result by number you originally thought", secondNum, prompt)
	reader.ReadString('\n')

	fmt.Println("Subtract ", subtraction, prompt)
	reader.ReadString('\n')
}

func main() {
	//way 1 to declare and assign
	// var firstNum int
	// firstNum = 10

	//way 2 to declare and assign
	// var secondNum = 20

	//way 3 to declare and assign value
	// thirdNum := 30

	//seed random num generator
	rand.Seed(time.Now().UnixNano())

	var firstNum = rand.Intn(8) + 2 //get num between 2 to 10
	var secondNum = rand.Intn(8) + 2
	var subtraction = rand.Intn(8) + 2
	var ans int

	getans(firstNum, secondNum, subtraction)

	ans = firstNum*secondNum - subtraction
	fmt.Println("Answer is ", ans)

}
