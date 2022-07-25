package main

import (
	"fmt"
	"runtime"
	"sync"
	"time"
)

func p1(label string) {
	fmt.Println("inside ", label)
	for i := 0; i < 5; i++ {
		fmt.Println(label, ":", i)
	}
	fmt.Println(label, " end")
}

func p11(label string, wg *sync.WaitGroup) {
	fmt.Println("inside ", label)
	for i := 0; i < 5; i++ {
		fmt.Println(label, ":", i)
	}
	wg.Done() //  == (*wg).Done()
	fmt.Println(label, " end")
}

func main() {
	fmt.Println("main execution started")
	fmt.Println("logical cores usable by current proces ", runtime.NumCPU())
	fmt.Println("num of goroutines that currently exists ", runtime.NumGoroutine())
	fmt.Println("OS ", runtime.GOOS)
	fmt.Println("arch ", runtime.GOARCH)

	// goroutine === function to be executed // can be executed concurrently, i.e. time-sharing
	// these functions are not directly mapped to a os thread
	// hence, go doesnt rely on os scheduler to schedule its tasks(lightthreads?) // go has its own scheduler
	// that go scheduler maps functions to os process/thread // m:n multiplexing or scheduling // m go routines to n processes
	// each goroutine will have a function to be executed and some call stack for variables and program counter
	// gorountine call stack is dynamic // not fixed sized static like os process of 1-2MBs // go stack can be as low as in KBs // as high in GBs
	// hence we can have 1000x such small task that can be executed  over 1 os thread

	// go scheduler // different from os scheduler
	fmt.Println("scheduler will schedule go codes on upto ", runtime.GOMAXPROCS(0), " os threads") // n in M:N multiplexing

	go p1("Job2")

	p1("Job1")
	// adding wait so that job2 is processed
	time.Sleep(time.Millisecond * 1)

	// waiting with wait groups // for letting the spawanned job complete
	var wg sync.WaitGroup
	wg.Add(2) // total jobs
	go p11("p11a", &wg)
	p11("p11b", &wg)

	wg.Wait()
	fmt.Println("main execution ended")

}
