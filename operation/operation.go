package main

import "fmt"

func main() {
	//math operation
	a := 3
	a *= a
	println(a, a/2)

	// comparisson operation
	b := 13
	fmt.Println(a == b)
	fmt.Println(a <= b)
	fmt.Println(a > b)

	fmt.Println(a == b && a > b)
	fmt.Println(a <= b || a > b)

	fmt.Println(a, b)
	a += b    // a = 9+13 = 22
	b = a - b //b = 22-13 = 9
	a = a - b //a = 22-9 = 13
	fmt.Println(a, b)

}
