package main

import (
	"fmt"
	"sync"
	// "time"
)

var a int = 0

func main() {
	/*
	go myFuncGo()
	myFunc()
	go myFuncGo()
	myFunc()

	time.Sleep(time.Second)
	*/

	var wg sync.WaitGroup
	go myFuncSync(&wg)
	wg.Add(1)

	myFunc()

	go myFuncSync(&wg)
	wg.Add(1)

	myFunc()

	wg.Wait()
}

func myFunc() {
	a += 1
	fmt.Println("prog: ", a)
}

func myFuncGo() {
	a += 1
	fmt.Println("go: ", a)
}

func myFuncSync(wg *sync.WaitGroup) {
	defer wg.Done()

	a += 1
	fmt.Println("go: ", a)
}
