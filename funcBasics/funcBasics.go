package main

import (
	"fmt"
	"log"
	"math"
	"os"
	"strings"
)

func f1(a int, b int) int {
	return a + b
}

func f12(a int, b int) (int, float64) {
	return a + b, math.Pow(float64(a), float64(b))
}

func f13(a int, b int) (result int) { // basically you dont need to init the values that will be returned // also dont need to return then as well, they will be auto returned // advisable to use in smaller function, not in big
	result = a + b
	return
}

func f14(a, b int) int {
	return a + b
}

func f15(a []int) {
	fmt.Printf(" %T,  %#v\n", a, a)
}

func f16(a ...int) { // so varargs variadic functons, try to create a slice of the variables sent // also if spreaded from an input slice, it creates a slice of the orignal slice only(sharing the backing array)
	fmt.Printf(" %T,  %#v\n", a, a)
}

func f17(a ...int) { // so varargs variadic functons, try to create a slice of the variables sent // also if spreaded from an input slice, it creates a slice of the orignal slice only(sharing the backing array)
	a[0] = 1000
}

func printAndSumAndProduct(msg string, nums ...float64) (float64, float64) { // when mixing nonvariadic args with variadic, keep them ahead of variadic ones //  TODO Can keep two variadic params???

	fmt.Println(msg)

	sum := 0.
	prod := 1.

	for _, v := range nums {
		sum += v
		prod *= v
	}

	return sum, prod

}

func getInfo(age int, names ...string) string {
	fullName := strings.Join(names, " ")
	returnString := fmt.Sprintf("Age : %d , FullName: %q", age, fullName)
	return returnString
}

func main() {

	// functions /// does some task
	// function in go cannot be overloaded
	// camelCase naming convention
	// can return multiple values
	// within same package, function names should be unique // obvio
	// main() int() reserved functions

	/*
		func (receiver) name(parameters) (returns) {
			//code -> function body here
		}

		TODO whats receiver???
	*/

	// simple function
	p1 := fmt.Println
	p1(f1(1, 2))

	// multiple returns
	_, r2 := f12(2, 3) // ignoring the additive return // using the pow return
	p1(r2)

	// named return
	p1(f13(1, 2))

	// variadic functions
	// that accepts the spread operator / ellipsis
	// the function that gets the variable args, gets it as a slice
	x := []int{1, 2, 4, 6}
	f15(x)
	f16(x...)
	f17(x...)
	p1("orignal backing array updated", x, " hence it created a slice on the same backing array")

	p1(printAndSumAndProduct("doing sum and prod", 2.1, 1., 4.))
	p1(getInfo(math.MaxInt, "albus", "percival", "wulfric", "brian", "dumbledore")) // basically amar <3

	// defer
	// this add to the stack of operataions that are to be performed once the function completes
	// defer function is executed after the function, but the arguments are evaluated then and there only
	// users :: for cleanup activites for open resources :) file closures streamclosures?
	file, err := os.Open("funcBasics.go")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()                   // will be executed once this function exec completes. and close the opened resources :)
	dataReadInBytes := make([]byte, 100) // create space to store data that will be read
	file.Read(dataReadInBytes)
	p1(string(dataReadInBytes))
	p1(strings.Repeat("##", 20))
	p1("1")
	defer p1("2")
	defer p1("3")
	p1("4") // order of print would be // 1 // 4 // 3 // 2

}
