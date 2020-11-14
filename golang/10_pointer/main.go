package main

import "fmt"

func main() {
	var i uint32 = 5

	// * = deferencing, & = address
	poi := &i
	fmt.Println(i, poi, *poi)

	*poi = 10
	fmt.Println(i, poi, *poi)

	var j uint32 = 20
	poi = &j
	fmt.Println(i, poi, *poi)

	x, y := 3, 5
	fmt.Println(x, y)
	swap(&x, &y)
	fmt.Println(x, y)
}

/*
// pass-by-value
func swap(a int, b int) {
	tmp := a
	a = b
	b = tmp
	fmt.Println(a, b)
}
*/

// pass-by-reference
func swap(a* int, b* int) {
	tmp := *a
	*a = *b
	*b = tmp
	fmt.Println(*a, *b)
}
