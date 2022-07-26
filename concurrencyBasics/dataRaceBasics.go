package main

import (
	"fmt"
	"sync"
)

func main() {

	// DATA RACE // unexpected update of the SHARED VARIABLE by concurrent tasks / goroutines

	// Some causes of "Data Race" in Go
	//https://eng.uber.com/data-race-patterns-in-go/

	// Checking Data Race // go can check for memoery that was accessed in a race // add the race flag while runnning the app
	// go run -race  dataRaceBasics.go

	// code with data race
	const rcount = 250
	var wg sync.WaitGroup
	wg.Add(2 * rcount)
	n := 0
	for i := 0; i < rcount; i++ {
		go func() {
			n++
			wg.Done()
		}() // anonymous function called without any param // when its executed, it will take whatever would be the current values of n and wg, at the moment in future when it gets cpu
		go func() {
			n--
			wg.Done() // 2nd issue here is that subtracts may happen before add
		}()
	}

	wg.Wait()
	fmt.Println(n)

}
