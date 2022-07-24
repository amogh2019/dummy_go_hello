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

type car struct {
	brand string
	cost  int
}

// function without method receiver
func modifyCar1(c car, newBrand string, newCost int) {
	c.brand = newBrand
	c.cost = newCost
}

// function with method receiver // type's instantiated variable  is passed as value
func (c car) modifyCar2(newBrand string, newCost int) {
	c.brand = newBrand
	c.cost = newCost
}

// function with method receiver // type's instantiated variable's pointer  is passed as value
func (c *car) modifyCar3(newBrand string, newCost int) {
	(*c).brand = newBrand
	(*c).cost = newCost
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

	// about method receivers of type pointers
	mycar := car{brand: "Audi", cost: 10000}
	p1(mycar)
	modifyCar1(mycar, "VW", 2000)
	p1(mycar, "still same, not updated")
	mycar.modifyCar2("VW", 2000)
	p1(mycar, "still same, not updated (values cloned and clone object updated, not original)")
	(&mycar).modifyCar3("VW", 2000) // way 1 of passing pointer to that variable
	p1(mycar, "orignal updated, since pointer is passed")
	mycar.modifyCar3("BMW", 4000) // way 2 of passing pointer to that variable // go does the translation to way 1 internally // ease for developers to pass pointers to modify type
	p1(mycar, "orignal updated, since pointer is passed")

	var mycarpointer *car
	mycarpointer = &mycar
	mycarpointer.modifyCar3("CAT", 2000) // way 3 of passing pointer to that variable // using a pointer already pointing to our variable
	p1(mycar, "orignal updated, since pointer is passed")

	// when to use pointers in method receivers?
	// 1. when you want to modify the values in initialized variables of the named type
	// 2. since everything is pass by value, we know that structs and things are copied. Use pointers in method receivers when you dont want large objects to be copied every time you call the methods and pass those big objects

	// also, method receivers are not allowed on namedTypes which are pointer types or interface type
	//type distance *int
	// func (d *distance) print() {} // compiler error

}
