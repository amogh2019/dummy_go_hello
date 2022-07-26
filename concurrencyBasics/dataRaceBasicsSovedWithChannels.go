package main

import (
	"fmt"
	"strings"
)

func main() {

	// explicit synchronization // threadsafety
	fmt.Println(strings.Repeat("####", 4))
	fmt.Println("writing race free access with channel ")

	const rcount = 250

	n := 0
	ch := make(chan int)
	for i := 0; i < rcount; i++ {
		go func(c chan int) {
			n++
			ch <- 1
		}(ch)
		fmt.Println("ith increment done", <-ch)
		go func(c chan int) {
			n--
			ch <- 1
		}(ch)
		fmt.Println("ith decrement done", <-ch)
	} // isn't this now almost sequential // you are spwanning go routine and waitng one by one

	fmt.Println(n)

}
