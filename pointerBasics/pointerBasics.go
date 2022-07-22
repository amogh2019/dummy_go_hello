package main

import (
	"fmt"
	"strings"
)

type Person struct {
	name string
	age  int64
}

func change(a *int) {
	*a = *a + 100
}

func incorrectChange(a int) {
	a = a + 100
}

func changeStruct(p *Person) {
	p.age = p.age + 1 // === (*p).age
	(*p).name += "hello"
}

func changeSlice(list []int) {
	for idx := range list {
		list[idx] += 10
	}
}

func changeMap(m map[string]int) {
	m["a"] = 10
	m["b"] = 20
	m["c"] = 30
}

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

	// WHY GO Pointers are easier / safe?
	// 1.
	// as said earlier
	// everything in go is passed by value
	// there is no reference // pass by references
	// the pointers that are passed, are actually passed by value
	// the passed pointers are clone of the orignal pointers, pointing to the same location

	/** 2.
	Memory location can still be accessed even after function has been called and scope in our opinion is destroyed
	Once a value is assigned to a pointer, with the exception of nil which I’ll cover in the next point, Go guarantees that the thing being pointed to will continue to be valid for the lifetime of the pointer. So

	func f() *int {
	        i := 1
	        return &i
	}

	is totally safe to do in Go. The compiler will arrange for the memory location holding the value of i to be valid after f() returns.
	*/

	// 3.
	// there is no pointer/array duality // like in c
	// so no pointer arithemetic
	// strings are not pointers in go, they are a type, hence always initialized to "" , and not null which is the most common cause of NullPointerException in Java / C

	fmt.Println(strings.Repeat("###", 8))

	// PASS BY VALUE for int strings  struct etc // dont modify arrays waise, instead pass slice
	x11 := 10
	p1("value of x is", x11)
	change(&x11)
	p1("value of x is", x11)
	incorrectChange(x11)
	p1("value of x is", x11)
	stru := Person{
		name: "gopher",
		age:  10,
	}
	p1(stru)
	changeStruct(&stru)
	p1(stru)

	// PASS BY VALUE for slices map
	// backing array/structure remains same
	// slice/map headers are cloned
	s1 := []int{1, 2, 3}
	p1(s1)
	changeSlice(s1)
	p1(s1)
	m1 := map[string]int{
		"a": 1,
		"b": 2,
	}
	p1(m1)
	changeMap(m1)
	p1(m1)

}
