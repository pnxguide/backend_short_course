package main

import "fmt"

type subtract func(a int, b int) int

func main() {
	fmt.Println(add(5, 6))
	fmt.Println(addShape(Rect{4, 3}, Rect{1, 2}))

	a := func() {
		fmt.Println("Anonymous Function")
	}

	a()

	func(str string) {
		fmt.Println("prog > ", str)
	}("something")

	var b subtract = func(a int, b int) int {
		return a - b
	}

	fmt.Println(b(4, 2))

	// higher-order function
	myHigherFunc(myFunc)

	c := createIncrement(5)
	fmt.Println(c(10))

	fact := factorial()
	for i := uint(0); i < 10; i++ {
		fmt.Println(fact(i))
	}
}

func myFunc() {
	fmt.Println("Hello!")
}

func myHigherFunc(someFunc func()) {
	someFunc()
	someFunc()
}

func createIncrement(n int) func(x int) int {
	return func(x int) int {
		return x + n
	}
}

// closure
func factorial() func(n uint) uint {
	var y uint = 1
	return func(n uint) uint {
		if n > 1 {
			y = n * y
		} else {
			return 1
		}
		return y
	}
}

/*
func myFunc(x int) {

}
*/

func add(a int, b int) int {
	return a + b
}

type Shape interface {
	Area() float64
	Perim() float64
}

type Rect struct {
	width, height float64
}

func (r Rect) Area() float64 {
	return r.width * r.height
}

func (r Rect) Perim() float64 {
	return 2*r.width + 2*r.height
}

func addShape(a Shape, b Shape) (float64, string) {
	return (a.Area() + b.Area()), "ok"
}
