package main

import (
	"fmt"
	"math"
)

type CentiMeters float64

func (c CentiMeters) isTooLong() {
	if c > 100 {
		fmt.Println("Toolong \n")
	} else {
		fmt.Println("Just Right \n")
	}
}

type rectangle struct {
	height float64
	width  float64
}

type shape interface {
	getArea() float64
	getPerimeter() float64
}

func (c rectangle) getArea() float64 {
	return c.width * c.height
}

func (c rectangle) getPerimeter() float64 {
	return 2*c.width + 2*c.height
}

func measureShape(s shape) {
	fmt.Printf("shape has an area of %f\n", s.getArea())
	fmt.Printf("shape has an perimeter of %f\n", s.getPerimeter())
}

type circle struct {
	radius float64
}

func (c circle) getArea() float64 {
	return math.Pi * c.radius * c.radius
}

func (c circle) getPerimeter() float64 {
	return 2 * math.Pi * c.radius
}

type square struct {
	width float64
}

func (s square) getArea() float64 {
	return s.width * s.width
}

func (s square) getPerimeter() float64 {
	return 4 * s.width
}
func main() {
	//myVar := CentiMeters(7.0)
	// fmt.Printf("Tipe %T and value %v \n", myVar, myVar)
	// myVar.isTooLong()
	myRectangle := rectangle{
		width:  20,
		height: 30,
	}
	myCircle := circle{
		radius: 5,
	}
	mySquare := square{
		width: 8,
	}
	fmt.Printf("Tipe %T and value %+v \n", myRectangle, myRectangle)
	measureShape(myRectangle)
	measureShape(myCircle)
	measureShape(mySquare)
}
