package main

import (
	"fmt"
	"time"
)

type names []string

func (n names) print() { // lowercase char at the start of function name

	for idx, v := range n {
		fmt.Println(idx, v)
	}
}

func main() {

	p1 := fmt.Println

	// receiver methods to named type / custom type
	// method == function that belongs to a type // usually function belogs to a package, but not here
	// method == function that takes receiver as an argument // the instance of the type on which this function is to be called, is the param

	// 1. use of an existing receiver method of an existing named type

	day := 24 * time.Hour //TODO how is this working? how are we able to do arithematic on different types? 24 is int and time.Hour is what, Duration?
	fmt.Printf("%T \n", day)
	seconds := day.Seconds() // using the receiver function of the named type Duration?
	fmt.Printf("%T \n", seconds)
	p1(seconds)

	// 2. custom receiver
	friends := names{"AA", "BB"}

	friends.print()      // 1st style of calling // calling as a method of the instantiated variable of the type
	names.print(friends) // 2nd style of calling // calling as a function of that type
}
