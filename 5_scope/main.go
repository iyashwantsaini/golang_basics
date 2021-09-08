package main

import (
	"scopeapp/packageone"

	"fmt"
)

func main() {
	var one = "One"
	fmt.Println(one)

	newString := packageone.PublicVar
	fmt.Println(newString)
	fmt.Println(packageone.ExportedOne)
}
