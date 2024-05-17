package main

/*
	Contains:
		Functions
		Structs
		RTTI
		Interfaces
*/
import (
	"fmt"
)

func greeting() string   { return "Hello World" }
func squareIt(i int) int { return i * i }

func demoRTTI(shp Shape) {
	switch v := shp.(type) {
	case Square:
		fmt.Printf("Found a square, the diagonal is %f\n", v.Diagonal())
		return
	case Circle:
		fmt.Printf("Is a Circle !!!, the diameter is %f\n", v.Diameter())
		return
	default:
		return
	}
}

func main() {
	fmt.Println(greeting())
	fmt.Println(squareIt(17))

	for i := 0; i < 10; i++ {
		if i%2 == 0 {
			fmt.Printf("Hello world for the %d-th time\n", i)
		}
	}

	fmt.Printf("Greeting: %s\n", greeting())

	//var square Square = Square{side: 1.0}
	//fmt.Printf("Area of Square %f\n", square.Area()) // Direct function call

	//var circle Circle = Circle{radius: 1.0}
	//fmt.Printf("Area of Circle %f\n", circle.Area()) // Direct function call

	// var shape1 Shape = Square{side: 2.0} // Huh ? YEAH.
	// var shape1 Shape = Circle{radius: 5.0}
	// fmt.Printf("Area of shape %f\n", shape1.Area())

	// It can also do run time type checking !!
	// demoRTTI(shape1)
}
