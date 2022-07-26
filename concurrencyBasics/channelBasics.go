package main

import "fmt"

func main() {

	// CHANNELS

	// currently viewing channels as queues (used to do sync b/w concurrent tasks)
	// the spawnned child will keep adding to queue
	// the main parent will try to make the queue empty before proceeding

	// used for inter goroutine communication
	// treat channel as a pointer passed in a function
	// channel three types // bidirectional chan // receiving chan<- // sender <-chan
	// channels has two operation // send operator <-  // receive operator

	//DECLARING CHANNEL
	var ch chan int // variable ch is a channel of type chan dealing with values of type int
	fmt.Println(ch) // nil

	//INIT CHANNEL
	ch = make(chan int)
	fmt.Println(ch)

	// SHORT DECLARATION
	c := make(chan int)

	go func() {
		// SEND ACTION
		c <- 4
	}() // if no other goroutine present // channel receiving illogical // hence error //  all goroutines are asleep - deadlock!

	//RECEIVE ACTION
	fmt.Println(<-c)

	close(c) // close to avoid further communication

	cpop := make(<-chan int)
	cpush := make(chan<- int) //TODO how are these used?
	// go func() {
	// 	cpush <- 4
	// }()
	// fmt.Println(<-cpop)
	_, _ = cpop, cpush

	ch1 := make(chan int)
	for ii := 1; ii < 7; ii++ {
		go func(i int, c chan int) {
			fact := 1
			for j := 2; j <= i; j++ {
				fact *= j
			}
			ch1 <- fact
		}(ii, ch1)

		//blocks and wait for channel to receive value
		fmt.Println("fact of ", ii, " is", <-ch1)
	}

}
