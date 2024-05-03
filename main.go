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

func average(current MovingAverage,
	result chan MovingAverage) {
	current.Average = float64(current.Total) / float64(current.Count) // Imagine that this is HEAVY lifting.
	result <- current                                                 // SEND MESSAGE
}
func display(current chan MovingAverage) {
	time.Sleep(time.Second)
	c := <-current //RECEIVE MESSAGE
	fmt.Println("CURRENT MOVING AVERAGE", c.Average)
}
func main() {
	speeds := []int{0, 20, 30, 40, 50, 60, 70}

	results := make(chan MovingAverage)
	var m = MovingAverage{}
	for i, speed := range speeds {
		m.Total = m.Total + speed
		m.Count = i + 1
		go average(m, results)
		display(results)
	}
	fmt.Println("END")
}
