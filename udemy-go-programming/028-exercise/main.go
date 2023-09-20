package main

import "fmt"

func main() {
	var i int8 = 127
	var j uint8 = 255

	fmt.Printf("The largest possible value for Type: %T is %v\n", i, i)
	fmt.Printf("The largest possible value for Type: %T is %v\n", j, j)
}


/*

write a program that declares two variables
● one variable to store a VALUE of TYPE int8
○ assign to it the largest number possible, then print it
● one variable to store a VALUE of TYPE uint8
○ assign to it the largest number possible, then print it

*/