package main

import "fmt"

func main() {
	var i int

	var l float64 = 10
	_ = 1

	fmt.Printf("var for zero value: %v\n", i)

	j := 2
	fmt.Printf("short declaration operator: %v\n", j)

	a, b := "kevin", 4
	fmt.Printf("multiple initializations: %v, %v\n", a, b)
	fmt.Printf("var when specificity is required: %v, type: %T\n", l, l)
}

/*

write a program that uses the following:
● var for zero value
● short declaration operator
● multiple initializations
● var when specificity is required
● blank identifier
print items as necessary to make the program interesting

*/