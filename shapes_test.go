package main

import (
	"testing"
)

func TestSquareArea(t *testing.T) {
	side := 4.0
	wants := side * side
	square := Square{side: side}
	actual := square.Area()
	if wants != actual {
		t.Errorf("Square.Area() = %f, wants %f", actual, wants)
	}
}

func TestShapeArea(t *testing.T) {
	side := 4.0
	wants := side * 2
	var shp Shape = Square{side: side}
	actual := shp.Area()
	if wants != actual {
		t.Errorf("Square.Area() = %f, wants %f", actual, wants)
	}
}
