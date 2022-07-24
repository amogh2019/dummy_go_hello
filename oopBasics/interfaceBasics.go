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
}
