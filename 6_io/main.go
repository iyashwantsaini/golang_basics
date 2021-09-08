package main

import (
	"fmt"
	"log"
	"strconv"

	"github.com/eiannone/keyboard"
)

func main() {
	err := keyboard.Open()
	if err != nil {
		log.Fatal(err)
	}

	defer func() {
		_ = keyboard.Close()
	}()

	fmt.Println("Press any key on the keyboard. Press ESC to quit.")

	coffees := make(map[int]string)
	coffees[1] = "Hot Coffee"
	coffees[2] = "Cold Coffee"
	coffees[3] = "Warm Coffee"
	coffees[4] = "Chilled Coffee"

	for i := 1; i < 5; i++ {
		fmt.Println(i, " : ", coffees[i])
	}

	for {
		char, key, err := keyboard.GetSingleKey()
		//error
		if err != nil {
			log.Fatal(err)
		}
		//break
		if char == 'Q' || char == 'q' {
			break
		}
		if key == keyboard.KeyEsc {
			break
		}
		//print
		if key != 0 {
			fmt.Println("You pressed", char, key)
		} else {
			x, _ := strconv.Atoi(string(char))
			fmt.Println("You pressed ", x)
		}
	}

	fmt.Println("Program exiting...")
}
