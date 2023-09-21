package main

import "fmt"

// interface to Person
type Person interface {
	ageTimesHeight() int
}

type person struct {
	heightInches int
	ageYears     int
}

func NewPerson(heightInches int, ageYears int) Person {
	return &person{
		heightInches: heightInches,
		ageYears:     ageYears,
	}
}

func (p person) ageTimesHeight() int {
	return p.ageYears * p.heightInches
}

func main() {
	var thisPerson interface{} = NewPerson(62, 71)

	switch thisPerson.(type) {
	case Person:
		fmt.Println("Person Interface")
	default:
		fmt.Println("Not a Person Interface")
	}

	fmt.Println(NewPerson(62, 71).ageTimesHeight())
}
