package main

import (
	"fmt"
	"os"
	"pymate"
)

func main() {
	fmt.Println(printThis("boop"))
}

func printThis(word string) string {
	dir, err := os.Getwd()
	stuff := fmt.Sprintf(dir, err, word)
	return stuff
}