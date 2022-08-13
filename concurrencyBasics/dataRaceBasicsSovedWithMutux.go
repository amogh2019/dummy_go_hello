package main

import (
	"fmt"
	"strings"
	"sync"
)

func main() {

	// explicit synchronization // threadsafety
	fmt.Println(strings.Repeat("####", 4))
	fmt.Println("writing race free access with mutux ")

	const rcount = 250
	var wg sync.WaitGroup
	wg.Add(2 * rcount)
	n := 0
	var m sync.Mutex // treat this like a pass to go inside the bank locker room
	for i := 0; i < rcount; i++ {
		go func() {
			m.Lock() // if a routine is at this line, it first tries to check that is m free to be taken and go inside locker room// if yes, take the pass // else wait
			n++
			m.Unlock() // leave the bank locker room and give the pass away
			wg.Done()
		}() // anonymous function called without any param // when its executed, it will take whatever would be the current values of n and wg, at the moment in future when it gets cpu
		go func() {
			m.Lock()
			defer m.Unlock() // will be queued once this function complete
			n--
			wg.Done()
		}()
	}

	wg.Wait()
	fmt.Println(n)

}
