package main

import "fmt"

func main() {

	p1 := fmt.Println
	// usual pointers
	// no pointer arithmetic like in c++

	// memory locations that store binary values
	// each location has an sequential address // address  // label for the memory location
	// a variable is a simple label for that address // a label for the memory location

	// &variable => address of the variable => pointer to that variable
	// *datatype => declarding a variable which will be a pointer
	// *pointer  => value at the address the pointer is pointing to    // TODO check when pointer = nil, what *pointer?

	// pointer can point to another pointer as well

	// creating pointer // 1 // pointing to existing address pointed by the variable
	z := 10
	zp := &z
	zpp := &zp
	p1(z)
	p1(zp, *zp)
	p1(zpp, *zpp, **zpp)

	// creating pointer // 2 // nil pointer
	var pp *int
	fmt.Printf("%#v\n", pp)

	// creating pointer // 3 // pointing to random new address, memory location
	p := new(int) // it creates a pointer called p that is a pointer to an int type
	x := 100
	fmt.Printf("%#v\n", p)
	p = &x // initializing p
	fmt.Printf("%#v ,  is point to x ka address %p \n", p, &x)

	// updating value in memory location
	*p = 40
	p1(x, "earlier 100")

}
