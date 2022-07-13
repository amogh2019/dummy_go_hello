package main

import (
	"fmt"
	"math"
	"os"
	"strconv"
)

func main() {

	// happy flow
	var input string = "34"
	i, err := strconv.Atoi(input)
	if err != nil {
		fmt.Println("for input ", input, "the error is", err)
	} else {
		fmt.Println("for input ", input, "the value of i is", i)
	}

	// error flow
	input = "34asdfadf"
	i, err = strconv.Atoi(input)
	if err != nil {
		fmt.Println("for input ", input, "the error is", err)
	} else {
		fmt.Println("for input ", input, "the value of i is", i)
	}

	// using simple if // this is like the for loop, initial statement first and then condition check, scope of variables remains
	input = "42"
	if i, err := strconv.Atoi(input); err != nil {
		fmt.Println("for input ", input, "the error is", err)
	} else {
		fmt.Println("for input ", input, "the value of i is", i)
	}
	input = "42a"
	if i, err := strconv.Atoi(input); err != nil {
		fmt.Println("for input ", input, "the error is", err)
	} else {
		fmt.Println("for input ", input, "the value of i is", i)
	}

	// another examples
	if args := os.Args; len(args) < 2 {
		fmt.Println("less than two args provided! please provide one more agrument", args)
	} else if val, err := strconv.Atoi(os.Args[1]); err != nil {
		fmt.Println("sorry, forgot to tell, provided 2nd argument must be an integer only", err)
	} else {
		fmt.Println("the square of the provided interger[", val, "] is", math.Pow(float64(val), 2))
	}
} 
