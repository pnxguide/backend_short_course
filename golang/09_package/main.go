package main

import (
	"fmt"
	"./utils"
)

func main() {
	a, b := 10, 2
	fmt.Println(utils.Add(a, b))
	fmt.Println(utils.Subtract(a, b))
	fmt.Println(utils.Multiply(a, b))
	fmt.Println(utils.Divide(a, b))
}
