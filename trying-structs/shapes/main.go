package main

import (
	"fmt"
	"math"
)

// define a shape interface type with a signle method to calculate area of the shape that several structs will satisfy
type Shape interface {
	GetArea() float64
}

// define an unexported recangle struct with necessary attributes to calculate area
type Rectangle struct {
	Height float64
	Width  float64
}

// make exported constructor function for Rectangle
func MakeRectangle(height float64, width float64) *Rectangle {
	return &Rectangle{
		Height: height,
		Width:  width,
	}
}

// add the GetArea function to satisfy the implementation requirements of the Shape interface
func (r Rectangle) GetArea() float64 {
	return r.Height * r.Width
}

// define an unexported circle struct with necessary attributes to calculate area
type Circle struct {
	Radius float64
}

// make exported constructor function for Circle
func MakeCircle(radius float64) *Circle {
	return &Circle{
		Radius: radius,
	}
}

// add the GetArea function to satisfy the implementation requirements of the Shape interface
func (c Circle) GetArea() float64 {
	return math.Pi * math.Pow(c.Radius, 2)
}

// define function to calculate total area of sevefral shapes with different formulas for area calculation
func TotalArea(shapes []Shape) float64 {
	var total float64
	for _, shape := range shapes {
		total += shape.GetArea()
	}
	return total
}

func main() {

	thisCircle := MakeCircle(13)          // constuct a cirle of radius 13
	thisRectangle := MakeRectangle(8, 17) // constuct a rectangle of heigh 8 and width 17

	circleArea := thisCircle.GetArea()       // calculate circle's area
	rectangleArea := thisRectangle.GetArea() // calculate rectangle's area

	// create slice of type Shape (the interface that implements both shape srtuct types) containing the circle and rectangle initialized above
	shapes := []Shape{
		thisCircle,
		thisRectangle,
	}

	totalArea := TotalArea(shapes) // calculate the trea acrifferen oss shape struc types

	fmt.Printf("A circle of type %T with radius %v has area %v\n", thisCircle, thisCircle.Radius, circleArea)
	fmt.Printf("A rectangle with height %v and width %v has area %v\n", thisRectangle.Height, thisRectangle.Width, rectangleArea)

	fmt.Printf("The total area of all shapes is %v\n", totalArea)
}
