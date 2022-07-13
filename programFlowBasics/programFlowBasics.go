package main

import (
	"fmt"
	"os"
	"strconv"
)

func main() {

	instock, price := true, 50

	// simple comparison
	if price > 100 {
		fmt.Println("Too Expensive!")
	}

	// simple comparison with and operator
	if instock && price < 100 {
		fmt.Println("Can buy!")
	}

	// this works in java but not in go // java would convert this to check if its null or not // but go wont, it need boolean check by developer only
	// if price {
	// 	fmt.Println("can proceed")
	// }
	if price > 0 {
		fmt.Println("can proceed to buy")
	}

	// command line args // go run programFlowBasics.go  arg1 arg2 hello world 123.45
	fmt.Println("these are the cmdline args :  \n", os.Args)
	fmt.Println("first argument is usually the path of the rile that is bring run :  \n", os.Args[0])

	// cmd args are always string // need to be parsed
	var k float64
	k, err := strconv.ParseFloat(os.Args[5], 64)
	_ = err
	fmt.Printf("the parsed float value of the 5th cmd line arg %f \n", k)
	fmt.Printf("type of the var from cmdline %T\n", os.Args[5])
	fmt.Printf("type of the var parsed %T\n", k)

}
