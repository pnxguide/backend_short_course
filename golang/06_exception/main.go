package main

import (
	"fmt"
	"log"
)

func main() {
	var myInt int
	fmt.Scanf("%d\n", &myInt)

	defer panicHandler()
	if myInt > 100 {
		panic("INTEGER_OVERFLOW")
	}
}

func panicHandler() {
	errMsg := recover()
	if errMsg == "INTEGER_OVERFLOW" {
		log.Println("Panic handled!")
	}
}
