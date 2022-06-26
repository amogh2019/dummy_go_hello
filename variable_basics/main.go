package main

import "fmt"

func main() {

	/*
		About the keyword var (About Variables)
	*/
	// full correct way // var is the keyword to compiler that
	//           , hey please assign a memoray location a label and treat it as variable
	var age int = 30
	fmt.Println(age)

	// short way // data type can be infered from the initiazed value
	var age1 = "pqr"
	fmt.Println(age1)

	// unused variables cause compile errors // mute them like with with caution
	var name = "pqe"
	_ = name // if this is not written // error is thrown

	// short delcaration // works only in blocked scope
	newName := "temp stringasdf"
	fmt.Println(newName)
	// age := 33    // this does not work on already defined variables // needs aleast one variable to be a new declaration

	// about multiple declaration

	// about multiple declaration // using short delcaration
	newName, newName1 := "new name", "new name 1"
	_, _ = newName, newName1 // leaving both blank // muting both // else just use them somewhere // print them maybe

	// about multiple declaration // using complete declaration
	// for better readablity
	// when we dont know initial values (else just use short )
	var (
		kk        int
		newName12 string = "hey"
		flag      bool
	)

	fmt.Println(kk, " ", newName12, " ", flag)

	// about multiple declaration with multiple assignments // using semi-complete declaration
	var ii, jj int
	ii, jj = 44, 55
	fmt.Println(ii, " ", jj)
	// swapping examples
	ii, jj = jj, ii
	fmt.Println(ii, " ", jj)

	var aa = 4
	var bb = 10.111

	// invalid type assignment
	fmt.Println("a ", aa, " b ", bb)
	// aa = bb this will break  // cannot assign float to int without casting
	aa = int(bb)
	fmt.Println("a ", aa, " b ", bb)

	var aa1 int
	// aa1 = "4"
	aa1 = 4
	_ = aa1

	var (
		intvar    int32
		floatvar  float64
		boolvar   bool
		stringvar string
	)
	fmt.Println(intvar, floatvar, stringvar, boolvar)

	// printing basics
	var (
		radius1 int    = 5
		radius2 int    = 155
		shape1  string = "circle"
	)

	const pi float64 = 3.14
	// pi = pi + 1

	/*
		About printing
	*/
	fmt.Printf("The diameter of %s of radius %+d is %f\n", shape1, radius1, float64(radius1)*2*pi)
	// using the generic format specififier
	fmt.Printf("The diameter of %v of radius %+v is %v\n", shape1, radius1, float64(radius1)*2*pi)
	// using quoted string
	fmt.Printf("The diameter of %q of radius %+d is %f\n", shape1, radius1, float64(radius1)*2*pi)
	// type of the variable
	fmt.Printf("figure is of type %T", shape1)
	// print boolean as true or false
	fmt.Printf("value of boolean var is %t  %v %t", boolvar, boolvar, boolvar)
	// print value in base 2 instaed of base 10
	fmt.Printf("the value of radius in base2 is %b\n", radius1)
	fmt.Printf("the value of radius in base2(8bit ans) is %08b\n", radius1)  // 8bits
	fmt.Printf("the value of radius in base2(16bit ans) is %16b\n", radius1) // 8bits
	// print value in base 16
	fmt.Printf("the value of radius in base16 is %x\n", radius2) // 8bits
	// formating decimal points
	fmt.Printf("value of pi (default decimal) %f \n", pi)
	fmt.Printf("value of pi (2 decimal decimal) %.3f \n", pi)

}
