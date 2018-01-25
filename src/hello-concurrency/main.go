package main

import (
	"fmt"
	"math"
	"sync"
)

func main() {
	names := []string{"phil", "noodles", "barbaro"}
	// declare new variable of the sync.WaitGroup data-type
	var wg sync.WaitGroup
	// len() returns the total num of names
	wg.Add(len(names))
	for _, name := range names {
		// execute each function call on a new go routine
		go printName(name, &wg)
	}
	// prevents program from exiting
	wg.Wait()
}

//must pass reference to WaitGroup so that we call done on the original value and not on a copy
func printName(n string, wg *sync.WaitGroup) {
	result := 0.0
	for i := 0; i < 10000000; i++ {
		result += math.Pi * math.Sin(float64(len(n)))
	}
	wg.Done()
	fmt.Println("name: ", n)
}
