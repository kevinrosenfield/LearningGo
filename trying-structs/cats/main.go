package main

import (
	"fmt"
	"math/rand"
)

type Cat interface {
	meow(volume int) (string)
}

type cat struct {
	color string
	age int
	cute bool
}
func CatMaker(color string, age int, cute bool) Cat {
	return &cat{
		color: color,
		age: age,
		cute: cute,
	}
}

func (c cat) meow(volume int) string {
	if volume < 5 {
		return "meow"
	}
	
	return "MEEEOWOOWWWOWOW"
}

func main() {
	c := CatMaker("blue", 13, false)
	r := rand.Intn(11)

	fmt.Println("Random volume:", r, "cat says: ", c.meow(r))
}