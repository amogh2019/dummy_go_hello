package main

import "fmt"

func increment(val int) func() int {

	return func() int {
		val++
		return val
	}
}

func main() {

	p1 := fmt.Println

	// Anonymous functions
	// create functions without name
	// aka closures
	// why need them? // so that we can pass functions as params to other function // we can return functions

	// Anonymous function 1 // invoking where its created
	func(name string) {
		p1("printing the name :", name)
	}("albus") // since we are not assiging to a variable to be called in future // its need to be call right here

	// Anonymous function 1 // getting it as a returned one
	incLambda := increment(10)
	fmt.Printf("%T\n", incLambda) // type:  func() int
	p1(incLambda())
}
