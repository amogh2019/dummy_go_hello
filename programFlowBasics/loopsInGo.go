package main

import (
	"fmt"
)

func main() {

	// same as in c/java
	// initialize part
	// condition part
	// some statement to run after each loop part
	// all are optional // hence it could result in initnite loop

	for i := 0; i < 10; i++ {
		fmt.Println("the value of i is", i)
	}

	// no init + conditioncheck + followupstatement
	// use of break (if not used, this just gives compile warning that future lines are unreachable, runtime goes in infinite loop)
	i := 0
	for {
		if i == 11 {
			break
		}
		fmt.Println("the value of A is", i)
		i++
	}

	// use of continue // with just condition in loop
	i = 0
	for i < 11 {

		if i < 5 {
			i++
			continue
		}
		fmt.Println("the value after 5 is", i)
		i++
	}

	// use of continue // 2
	i = 0
	for ; i < 11; i++ {

		if i < 5 {
			continue
		}
		fmt.Println("the value after 5 is", i)
	}

	// there is no do-while loop in GO
}
