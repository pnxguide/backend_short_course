package main

import (
	"fmt"
	/*
	"bufio"
	"log"
	"os"
	"strconv"
	"strings"
	*/
)

func main() {
	/*
	reader := bufio.NewReader(os.Stdin)
	tmp, _ := reader.ReadString('\n')
	tmp = strings.TrimSuffix(tmp, "\n")

	// ParseInt(val, base, size)
	myInt64, err := strconv.ParseInt(tmp, 10, 32)
	if err != nil {
		log.Fatal("error!")
	}

	myInt := int(myInt64)
	fmt.Println(myInt)
	*/

	var myInt int
	fmt.Scanf("%d\n", &myInt)	// no buffer usage, slow! (but easy for C dev)

	for i := 0; i < myInt; i++ {
		for j := 0; j < myInt; j++ {
			fmt.Print("*")
		}
		fmt.Println()
	}
}
