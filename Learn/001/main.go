package main

import "fmt"

func main() {

	//  not work because -> e is declared and not used
	/*
		b, c, d, e := 0, 1, 2, 3
		fmt.Println(b, c, d)
	*/

	// will work
	/*
		a := 25
		fmt.Println((a))

		b, c, d, _, f := 0, 11, 22, 33, "Ali Hamood"
		fmt.Println(b, c, d, f)
	*/
	var g int
	fmt.Println(g)

	//initial value of g will be g = 0
	var m int = g + 1
	fmt.Println("m is = ", m)

	g = 43
	fmt.Println(g)

	// declare a variable to hold a VALUE of a certain TYPE
	// initializing a variable
	var h int = 44
	fmt.Println(h)
}
