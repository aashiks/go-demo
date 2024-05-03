package main

/*
	Contains:
		Go routine demo
*/
import (
	"fmt"
	"time"
)

type MovingAverage struct {
	Total   int
	Count   int
	Average float64
}

func sum(speeds []int,
	results chan MovingAverage) {
	var m = MovingAverage{}
	for i, speed := range speeds {
		m.Total = m.Total + speed
		m.Count = i + 1
		m.Average = float64(m.Total) / float64(m.Count)
		results <- m
	}
}
func display(current chan MovingAverage) {
	
	c := <-current
	fmt.Println("CURRENT MOVING AVERAGE", c.Average)
}
func main() {
	speeds := []int{0, 20, 30, 40, 50, 60, 70}
	results := make(chan MovingAverage, 10)
	go sum(speeds, results)
	time.Sleep(time.Second)
	go display(results)
	fmt.Println("END")
}
