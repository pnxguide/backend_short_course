package main

import (
	"fmt"
	"math"
)

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

type Circle struct {
	radius float64
}

func (c Circle) Area() float64 {
	return math.Pi * c.radius * c.radius
}

func (c Circle) Perim() float64 {
	return 2 * math.Pi * c.radius
}

func measure(s Shape) {
	fmt.Println(s.Area(), s.Perim())
}

func main() {
	var expVar int32 = 5
	fmt.Printf("%d\n", expVar)

	// shorthand declaration
	impVar := "Hello"
	fmt.Printf("%s\n", impVar)

	var a1, a2 byte
	a1 = 'a'
	a2 = 'b'
	fmt.Printf("%c %c\n", a1, a2)

	var (
		var1 string
		var2 byte
	)

	// multiple-line string
	var1 = `
		Hello World!
	`
	var2 = 65
	fmt.Printf("%s %c\n", var1, var2)

	var infVar interface{}

	// type assertion
	s, ok := infVar.(string)
	fmt.Println(s)
	fmt.Println(ok)

	infVar = "Hello"
	s, ok = infVar.(string)
	fmt.Println(s)
	fmt.Println(ok)

	myRect := Rect{width: 4, height: 6}
	myCircle := Circle{5}
	measure(myRect)
	measure(myCircle)
}
