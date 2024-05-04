package main

import "fmt"

func main() {
	// adams := 42
	// fmt.Printf("42 as binary, %b \n", adams)
	// fmt.Printf("42 as hexadecimal, %x \n", adams)

	// print these values as both binary & hexadecimal
	a, b, c, d, e, f := 5, 11, 22, 33, 44, 55

	/*
		\t for tap space
		\n for new line
	*/
	fmt.Printf("%v \t %b \t %x \n", a, a, a)
	fmt.Printf("%v \t %b \t %x \n", b, b, b)
	fmt.Printf("%v \t %b \t %x \n", c, c, c)
	fmt.Printf("%v \t %b \t %x \n", d, d, d)
	fmt.Printf("%v \t %b \t %x \n", e, e, e)
	fmt.Printf("%v \t %b \t %x \n", f, f, f)
}
