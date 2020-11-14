package main

import "fmt"

func main() {
	fmt.Println("Hello, World!")
}

func init() {
	fmt.Println("Init! 1")
}

func init() {
	fmt.Println("Init! 2")
}