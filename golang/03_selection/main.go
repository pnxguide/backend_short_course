package main

import (
	"bufio"
	"fmt"
	// "log"
	"os"
	"strings"
)

func main() {
	reader := bufio.NewReader(os.Stdin)

	myStr, _ := reader.ReadString('\n')

	/*
	myStr, err := reader.ReadString('\n')
	if err != nil {
		log.Fatal(err)
	}
	*/

	myStr = strings.TrimSuffix(myStr, "\n")

	fmt.Println(myStr)

	if myStr == "Hi!" {
		fmt.Println("Hello!")
	} else {
		fmt.Println("Bye!")
	}

	switch myStr {
	case "A":
		fmt.Println("AA")
	case "B":
		fmt.Println("BB")
		fallthrough
	case "C":
		fmt.Println("CC")
	}
}
