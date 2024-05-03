package main

/*
	Contains:
		Go routine demo
*/
import (
	"fmt"
)

func greeting(from string) {
	for i := 0; i < 3; i++ {
		fmt.Println(from, ":", i)
	}
}

func main() {
	greeting("main")
	go greeting("go routine") //non blocking, light weight thread
	//time.Sleep(time.Second)   //: What happens here ?
	fmt.Println("END")
}
