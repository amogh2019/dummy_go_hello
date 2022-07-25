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

type threeDShape interface {
	volume() float64
}
type cube struct {
	side  float64
	color string
}

func (c cube) area() float64 {
	return math.Pow(float64(c.side), 2) * 6
}
func (c cube) peri() float64 {
	return 12 * c.side
}
func (c cube) volume() float64 {
	return math.Pow(float64(c.side), 3)
}

type geometry interface {
	twoDShape
	threeDShape
	getColor() string
}

func (c cube) getColor() string {
	return c.color
}
func printDetailsGeo(g geometry) {
	fmt.Println(g)
	fmt.Println("area", g.area())
	fmt.Println("peri", g.peri())
	fmt.Println("vol", g.volume())
	fmt.Println("color", g.getColor())
}

// interface with no method // is inherited by all types
type empty interface {
}
type person struct {
	info        empty       // can store any value
	anotherinfo interface{} // can store any value
	name        string
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
	// since interface type variables // can have dynamic types (of the mapped object) // we need to first cast it specifically, and then use  the desired dynamic method
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

	// interface nesting
	// interface cannot inherit multiple interface
	// need to create a new interface  // spefic the child interfaces(this brings all their methods) with other new methods // include, embedded
	chotaCube := cube{side: 7, color: "yello"}
	fmt.Println(chotaCube)
	var geoKaShape geometry
	geoKaShape = chotaCube
	printDetailsGeo(geoKaShape)

	// empty interface
	// can be used as a representation for any type
	// hence be careful in using them // since this bypassed strict typechecking // code will break at runtime only and not on compile time
	// used in functions // where we need to pass variables // and type is unknown, or not determined // like using Object in java
	var emptyVar1 empty
	_ = emptyVar1
	var emptyVar interface{}
	emptyVar = "stting"
	fmt.Println(emptyVar)
	emptyVar = 100
	fmt.Println(emptyVar)
	emptyVar = []int{1, 2, 3}
	fmt.Println(emptyVar)
	// fmt.Println(len(emptyVar)) // len is not a method on empty interface // but on dynamic instantiated object slice // so cast to slice first
	fmt.Println(len(emptyVar.([]int)))

	p := person{name: "yolo"}
	p.info = "random info" // assigning a string
	p.anotherinfo = 100
	fmt.Println(p)
	p.info = []float64{2.2, 3.4, 5.4, 44} // assiging a slice
	fmt.Println(p)
}
