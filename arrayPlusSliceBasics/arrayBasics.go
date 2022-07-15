package main

import (
	"fmt"
	"strings"
)

func main() {

	// arrays are not vectors // they are static and len is known at compile time only
	// arrays are not like List<> pointers // each array, since known at compile time only, is a new initialization // if someelse other array connects to it, its bascially creating a copy of the values, it actually doesnt point

	// simple declaration with 0
	var numbers [4]int // var is keyword // numbers is variable name // type of the variable number is :  [4]int  means arrayOfSize4OfIntegerType

	fmt.Printf("%v \n", numbers)
	fmt.Printf("%#v \n", numbers)

	// equals declaration
	var a1 = [4]float64{}
	fmt.Printf("%#v \n", a1)

	// another declaration with init values
	a2 := [4]int{1, 2, 3, 4}
	fmt.Printf("%#v \n", a2)

	// partial declaration with init values
	a3 := [4]int{1, -2}
	fmt.Printf("%#v \n", a3)

	// init value at particular index
	a31 := [10]int{3: 1, -2, 7: 222, 0: 999}
	fmt.Printf("%#v \n", a31)

	// ellipsis operator when you want complier to determine the size itself, when many values declared
	a4 := [...]int{1, -2, 2, 4, 2, 5, 2, 4}
	fmt.Printf("%#v \n", a4)

	// multiline declaration // *** needs a comma at the last element as well ***
	a6 := [...]int{
		1,
		-2,
		2,
		4,
		2, // if curly brace is not placed here // it needs the comma to be present
	}

	fmt.Printf("%#v \n", a6)

	// since array is known compiler time // array out of bounds is thrown at compiler time only
	a7 := [...]int{1, 2, 5}
	// fmt.Println(a7[10])
	fmt.Println(a7[2])

	// iteration array 1 /// range returns index, and the value
	for idx, val := range a7 {
		fmt.Printf("at idx %d , val is %d\n", idx, val)
	}
	fmt.Println(strings.Repeat("###", 8)) // just to mark segregation on console // beautification
	// iteration array 2 /// usual
	for i := 0; i < len(a7); i++ {
		fmt.Printf("at idx %d , val is %d\n", i, a7[i])
	}

	// 1d
	var a8 = [...]int{1, 3, 4}
	fmt.Println(a8)

	// 2d
	a9 := [2][3]int{
		{1, 2, 3}, // [3]int can be ignored here
		[3]int{1, 2, 3},
	}
	fmt.Println(a9)

	// cloning array is basically deep cloning // creates new array only
	a10 := a8
	// equating is values match and order match
	fmt.Println("is a8 equals a10", a8 == a10)
	a10[1] = 11212
	// since value changed and they both are different initializations, there will be mismstach /// had it been slices, it would have chnaged in both and they would be equal // since slices (would be something like pointers to same memory)
	fmt.Println("is a8 equals a10", a8 == a10)

}
