package main

import "fmt"

var (
	s string 	= "woop"
	i int 		= 98786
	f float64 	= 234.325
)


func main() {

	fmt.Printf("Value: %v, Type: %T\n", s, s)
	fmt.Printf("Value: %v, Type: %T\n", i, i)
	fmt.Printf("Value: %v, Type: %T\n", f, f)

}




/*

write a program that uses the following:
● for a VARIABLE storing a VALUE of TYPE
○ string
○ int
○ float64
● use print verbs to show
○ value
○ type

*/