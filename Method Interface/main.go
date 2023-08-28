package main

import (
	"fmt"
	"math"
)

type CentiMeters float32

type rectangle struct {
	width  float32
	height float32
}

type circle struct {
	radius float32
}

type square struct {
	width float32
}

type shape interface {
	getArea() float32
	getPerimeter() float32
}

func (c CentiMeters) IsTooLong() {
	if c > 100 {
		fmt.Printf("Too Long!\n")
	} else {
		fmt.Printf("Just right!\n")
	}
}

func (r rectangle) getArea() float32 {
	return r.width * r.height
}

func (r rectangle) getPerimeter() float32 {
	return 2*r.width + 2*r.height
}

func (c circle) getArea() float32 {
	return math.Phi * c.radius * c.radius
}

func (c circle) getPerimeter() float32 {
	return 2 * math.Phi * c.radius
}

func (s square) getArea() float32 {
	return s.width * s.width
}

func (s square) getPerimeter() float32 {
	return 4 * s.width
}

func measureShape(s shape) {
	fmt.Printf("Shape has an Area of %f\n", s.getArea())
	fmt.Printf("Shape has an Perimeter of %f\n", s.getPerimeter())
}

func main() {
	myVar := CentiMeters(101)
	fmt.Printf("Type: %T, value %v\n", myVar, myVar)
	myVar.IsTooLong()

	myRec := rectangle{
		width:  30,
		height: 32,
	}

	myCir := circle{
		radius: 7,
	}

	mySqu := square{
		width: 5,
	}
	fmt.Printf("Type: %T, value %+v\n", myRec, myRec)

	fmt.Printf("%.2f\n", myRec.getArea())
	measureShape(myRec)
	measureShape(myCir)
	measureShape(mySqu)
}
