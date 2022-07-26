package main

import (
	"fmt"
	"time"
)

func main() {

	start := time.Now()

	ch1 := make(chan string)
	ch2 := make(chan string)

	go func() {
		fmt.Println("starting 1 at", time.Since(start))
		time.Sleep(2 * time.Second)
		ch1 <- "Hi from 1"
	}()
	go func() {
		fmt.Println("starting 2 at", time.Since(start))
		time.Sleep(2 * time.Second)
		ch2 <- "Hi from 2"
	}()

	fmt.Println("waiting to get messages")
	//SELECT is like switch but for channels receviers
	// used when we want to wait on multiple channels receivers or goroutines simualtaneously
	for i := 0; i < 2; i++ {
		select {
		case msg1 := <-ch1:
			fmt.Println(msg1)
		case msg2 := <-ch2:
			fmt.Println(msg2)
			// default:     // if we uncomment this, select will never wait for receivers if they are empty // for this to work // introduce sleep earlier to make main leave cpu, and a routine to get cpu and produce a message in channel // if channel not empty case will work
			// 	fmt.Println("no message in any channel yet")
		}
	}

	fmt.Println("lapsed ", time.Since(start)) // will take around two seconds // main queues goroutines // when main is blocked to receive // routine 1 activates // routine 1 sleeps for 2s // routine 2 activets and sleeps for 2 s // main gets 1 values at around 2 seconds and then the next one also immediately after /// overall blocked time is the time spend in gorountines sleep
}
