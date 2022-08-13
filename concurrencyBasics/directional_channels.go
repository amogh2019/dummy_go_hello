package main

import (
	"log"
	"strconv"
	"time"
)

type res struct {
	i   int
	val int64
}

func main() {

	// This is to show an example of
	// 	1. Use of push / pop  channels
	//  2. Advantages of using concurrency

	start := time.Now()
	orchestrate(1)
	spentWith1 := time.Since(start)

	start = time.Now()
	orchestrate(4)
	spentWith4 := time.Since(start)
	log.Println("Time Spent with 1 worker", spentWith1)
	log.Println("Time Spent with 4 worker", spentWith4)

	/*
		sample output
		2022/08/13 14:38:10 Time Spent with 1 worker 15.8624282s
		2022/08/13 14:38:10 Time Spent with 4 worker 7.701159842s
	*/

}

func orchestrate(noOfWorkers int) {

	jobs := make(chan int, 200) // making this a buffered channel to keep sending byorchestrator and receiving by worked async
	results := make(chan res, 200)

	for i := 0; i < noOfWorkers; i++ {
		go worker(jobs, results, strconv.Itoa(i)) // starting the worked to listen on channel
	}

	for i := 0; i < 46; i++ {
		jobs <- i // sending jobs to channel
	}
	close(jobs)

	for i := 0; i < 46; i++ { // here i am not relying on result channel close // instead relying that 50 messages were expected // since a workedroutine1 may take lot of time to process last message and the other workedroutine2 may see now that jobs channel is closed, and so may send a close signal even before result of last message was sent
		processedValue := <-results
		log.Println("ith", strconv.Itoa(i), " result at orchestrator : ", processedValue)
	}
	defer close(results)

}

// this will keep popping a message from the jobs channel, execute a long running processingFunction, push the result in results channel
func worker(jobs <-chan int, results chan<- res, workerId string) {

	for {
		ith, isOpen := <-jobs
		if isOpen {
			results <- res{
				i:   ith,
				val: fib(ith),
			}
		} else { // if a channel is closed // whenever any worker asks for message any time in future // will get channel closed
			log.Println("closing worker[", workerId, "]")
			break
		}
	}
}

// an intended slower implementation of finding nth fibonacci
func fib(i int) int64 {
	if i < 0 {
		return -1 // to denote invalid for now
	}
	if i < 2 {
		return int64(i)
	}
	return fib(i-1) + fib(i-2)
}
