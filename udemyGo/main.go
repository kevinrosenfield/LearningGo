package main

import "fmt"

func main() {
	a, b, c, d, e, f := 0, 1, 2, 3, 4, 5

	fmt.Printf("%v \t binary: %b \t hex: %#X\n", a, a, a)
	fmt.Printf("%v \t binary: %b \t hex: %#X\n", b, b, b)
	fmt.Printf("%v \t binary: %b \t hex: %#X\n", c, c, c)
	fmt.Printf("%v \t binary: %b \t hex: %#X\n", d, d, d)
	fmt.Printf("%v \t binary: %b \t hex: %#X\n", e, e, e)
	fmt.Printf("%v \t binary: %b \t hex: %#X", f, f, f)
}