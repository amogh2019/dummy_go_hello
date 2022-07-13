package main

import (
	"fmt"
)

func main() {

	// labels  // can be used by break / continue / goto
	// break + label // usually used to bring execution to the outer of a nested for loop
	// label names dont conflict with other var/constant/function names // they are maintained elsewhere
	// if defined, will have to use it somewehere
	// scope // scope of the body where defined // not in nested body??

	names := [5]string{"A", "B", "C", "D", "E"}
	blacklistedNames := [5]string{"C", "D", "E"}

	// this only breaks inner loop // we will have to write extra code to break the outer loop as well // by keeping flags and what not
	for idx, name := range names {
		for idx2, blackListedName := range blacklistedNames {
			_ = idx2
			if name == blackListedName {
				fmt.Println("A the names", names, " have atleast one blacklisted name", blackListedName, "at index", idx)
				break
			}
		}
	}

	// this only breaks both loops //

outerLabel:
	for idx, name := range names {
		for idx2, blackListedName := range blacklistedNames {
			_ = idx2
			if name == blackListedName {
				fmt.Println("B the names", names, " have atleast one blacklisted name", blackListedName, "at index", idx)
				break outerLabel
			}
		}
	}
	fmt.Println("This line is after both the loops!")

	// goto + label // works as a loop
	// in other language // goto can refer a label anywherer // but here it can refer to labels as long has no new var are initialized from point of execution till the label in future?

	i := 1
loopLabel:
	if i < 5 {
		fmt.Println("val of i is", i)
		i++
		goto loopLabel
	}
	fmt.Println("This line is after the gotoloop")

	//  goto todo //ERROR it's not permitted to jump over the declaration of x
	//  x := 5
	// todo:
	//  fmt.Println("something here")

}
