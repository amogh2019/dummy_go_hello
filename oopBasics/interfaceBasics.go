package main

import (
	"fmt"
	"math"
)

type rect struct {
	len, bdt float64
}
type circle struct {
	rad float64
}

func (r rect) area() float64 {
	return r.bdt * r.len
}
func (c circle) area() float64 {
	return math.Pi * math.Pow(c.rad, 2)
}
func (r rect) peri() float64 {
	return r.bdt + r.len
}
func (c circle) peri() float64 {
	return 2 * math.Pi * c.rad
}

// life without interface // had to write same print method for all  shapes :(
func printDetailsCircle(c circle) {
	fmt.Println(c)
	fmt.Println(c.area())
	fmt.Println(c.peri())
}
func printDetailsRect(c circle) {
	fmt.Println(c)
	fmt.Println(c.area())
	fmt.Println(c.peri())
}

// life with interface
type twoDShape interface {
	area() float64
	peri() float64
}

// since both circle and rect named types already implement as methods(receiverfunctions) the functions that are defined in the interface
// go automatically says that twoDShape interface is implemented by circle and rect // we dont need to explicitely specifcy

func printDetails(s twoDShape) {
	fmt.Println(s)
	fmt.Println(s.area())
	fmt.Println(s.peri())
}

type equilateralTri struct {
	len float64
}

func (t equilateralTri) peri() float64 {
	return 3 * t.len
}

func (c circle) volume() float64 {
	return math.Pow(c.rad, 3)
}

func main() {

	c := circle{rad: 5.}
	r := rect{len: 2., bdt: 3.}

	// since both rect and circle named types, implement both the methods of the interface twoDShape, // hence we can call the functions which require an instance of twoDShape
	printDetails(c)
	printDetails(r)

	t := equilateralTri{len: 4.}
	// we cannot call printDetails() of functions that require twoDShape, for equilateralTri, since it does not implement all the methods of that interface
	// printDetails(t) // cannot use t (variable of type equilateralTri) as twoDShape value in argument to printDetails: equilateralTri does not implement twoDShape (missing method area)compilerInvalidIfaceAssign
	fmt.Println(t.peri())

	// nil interface
	var blank twoDShape
	fmt.Printf("%T %#v\n", blank, blank)
	// printDetails(blank) // runtime errors  // nil pointer // no compile time issus
	ball := circle{rad: 3.}
	blank = ball
	fmt.Printf("%T %#v\n", blank, blank)
	printDetails(blank) // will work now, mapped to concrete object with underlying type circle // its dynamic type is now circle
	// interface type //works such that // it holds the reference to its underlying  concrete type
	// variables's type in go doesnot change // but interface in go can have dynamic types // kal nil, aaj circle, parso rect

	// interface type assertion
	fmt.Println(ball.volume()) // works since ball is of type circle // and volumne is methodreceiver for type circle
	// blank.volume() // wont work // even though dynamic type is circle
	// since interface type variables // can have dynamic types (of the mapped object) // we may need some type checking in code
	fmt.Println(blank.(circle).volume()) //this works now // we say // work only if the type is circle
	castedBlank, ok := blank.(circle)
	if ok == true {
		fmt.Println(castedBlank.volume())
	}
	switch tt := blank.(type) {
	case circle:
		fmt.Println("interface is of mapped type circle", tt)
	case rect:
		fmt.Println("interface is of mapped type rect", tt)
	}

}
