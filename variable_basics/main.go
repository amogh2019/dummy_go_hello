package main

import "fmt"

func main() {

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
}
