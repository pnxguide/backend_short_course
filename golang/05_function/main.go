package main

import "fmt"

func main() {
	fmt.Println(add(5, 6))
	fmt.Println(addShape(Rect{4, 3}, Rect{1, 2}))
}

func myFunc() {
	fmt.Println("Hello!")
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
	return 2 * r.width + 2 * r.height
}

func addShape(a Shape, b Shape) float64 {
	return a.Area() + b.Area()
}
