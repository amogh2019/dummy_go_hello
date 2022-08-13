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
	// FIFO 
	// async nature 
	// When to use? when we dont want to block sender till a msg is consumer // eg. when rate of consumption is slightly slower than production // so not to block message production, that is, it gives a small relief to the sender, since usually sender is blocked till some other routine consumes the message like in an unbuffered, but here, sender is not blocked and continues
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
	} 
	// How did the receiver know that sender sent just 6 and then sender stopped receiving from channel?
	// 	Ans. since sender is sending a close() signal,
	//		1. when receiver receives in this manner : <-ch  it returns {value, isOpen}, receiver can check on isOpen flag
	//		2. when you receive using range // loop automatically stops when isOpen is false
	// TODO what will happen and receiving is less ??

	time.Sleep(time.Second * 2) // giving time for go routine to complete and print its last line

	// receiving after channel closed // non blocking // default value of that type
	fmt.Println(<-bch) // prints empty string
	close(ch)
	fmt.Println(<-ch)
}
