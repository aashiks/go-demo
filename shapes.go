package main

import "math"

type Shape interface {
	Area() float64
}

type Square struct { // Note: no reference to the interface AT ALL
	side float64
}

// Note the function declaration.
func (sq Square) Area() float64     { return sq.side * sq.side }
func (sq Square) Diagonal() float64 { return sq.side * math.Sqrt2 }

type Circle struct { // Nothing here too.
	radius float64
}

func (c Circle) Area() float64     { return math.Pi * math.Pow(c.radius, 2) }
func (c Circle) Diameter() float64 { return 2 * c.radius }
