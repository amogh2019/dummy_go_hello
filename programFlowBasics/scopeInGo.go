package main

import (
	"fmt" // this is file scoped // means that even if someone imports this package(class) // user will still have to import fmt in its package, and this fmt will not be accessable by that user
)

const AA int = 10 // can be used by any function belonging to this package
var aaa int = 12  // package scope variables need not to be muted if not used // since we expect that any future function in this package might use it

func main() {

	const AA int = 11 // local / block scoped  // redeclared since there is also a package scoped one
	fmt.Println("local scope value", AA)

	f1()
}

func f1() {
	fmt.Println("packge scope value", AA)
}
