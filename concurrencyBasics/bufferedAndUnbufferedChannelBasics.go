package main

import (
	"fmt"
	"strconv"
	"time"
)

func main() {

	// UNBUFFERED CHANNEL
	// this channel has no space to hold temp value, hence
	// sender will be blocked, till there is an actual receiver ready to receive
	// receiver will be blocked, till there is a sender ready to send
	// also called sync channels
	ch := make(chan string)
	go func(c chan string) {
		fmt.Println("inside go routine : start")
		c <- "[HI FROM GO ROUTINE]" // this channel will be blocked until the receiving main channel completes it sleep and is ready to receive
		fmt.Println("inside go routine : finished")
	}(ch)

	fmt.Println("inside main routine : about to sleep")
	time.Sleep(time.Second * 2)
	fmt.Println("inside main routine : sleep over")

	fmt.Println("inside main routine : get value : start")
	val := <-ch
	fmt.Println("inside main routine : get value : end :", val)

	time.Sleep(time.Second * 2) // giving time for go routine to complete and print its last line

	// BUFFERED CHANNEL
	// this channel has space to hold temp value, hence
	// sender will send, till there is space to hold value
	// sender will be blocked, if no space left, till atleast something is consumed
	// receiver will be blocked, till all buffer space empty
	// receiver will work otherwise
	// FIFO // async  nature
	// When to use? when rate of consumption is slightly slower than production // so not to block production??
	bch := make(chan string, 3) // channel with buffer size
	go func(c chan string) {

		for i := 0; i < 6; i++ {
			fmt.Println("inside go routine : start", i+1)
			c <- "[HI FROM GO ROUTINE]" + strconv.Itoa(i+1) // this channel will be blocked until the receiving main channel completes it sleep and is ready to receive
			fmt.Println("inside go routine : finished", i+1)
		}
		close(c) // not always required // sometimes helpful for receivers to know now that no more things to receive

	}(bch)

	fmt.Println("inside main routine : about to sleep")
	time.Sleep(time.Second * 2) // to give  go routine some time to start // befroe the main finishes
	fmt.Println("inside main routine : sleep over")

	fmt.Println("inside main routine : get value : start")
	for vall := range bch { // == val := <-bch
		fmt.Println("inside main routine : get value : end :", vall)
	} //TODO what will happen // is sending is more in channed // and receiving is less ??  or vice versa

	time.Sleep(time.Second * 2) // giving time for go routine to complete and print its last line

	// receiving after channel closed // non blocking // default value of that type
	fmt.Println(<-bch) // prints empty string
	close(ch)
	fmt.Println(<-ch)
}
