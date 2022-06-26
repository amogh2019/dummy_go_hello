package main

import (
	"fmt"
)

func main() {

	/*
		Why constants? // Bruh these are like ENUMS // labels that are resolved at compile time
		you could use variables, but if we know that some variables will have fixed values, we can let Go know it already,
		hence, it can help detects error on compiler time only // rather than actual runtime
		not even go // even ides like vscode can also know // ki bruh since its a const // i know its value already // lemme just check it its used correctly, on your behalf :)
	*/

	const c1 int = 7
	// unlike variable // you need not use it  // and need not mute it

	const c2 = 14

	const c3, c4 = 0, 10
	fmt.Println(c1, c2, c3, c4)

	// initial value gets translated to all contants // omitting type and value in const means repeating the type and value
	const (
		c11 = 100
		c12
		c13
	)
	fmt.Println(c11, c12, c13)

	// helps in early error detectiong
	k1 := 10
	k2 := 0
	const c21 = 10
	const c22 = 0
	// variable divide by 0 // only delected at runtime // (that is a production bug :/ )
	// fmt.Println(k1 / k2)
	// this breaks at compile time only (that is in developemnt only and the developer can fix right awat)
	// fmt.Println(c21 / c22)
	_, _ = k1, k2

	/*
		Rules
	*/
	//1. cannot modify constnat
	// c1 = 14

	//2. cannot assign dynamic runtime values // values of functions or variables which run on runtime
	// const c33 = k1
	// const c33 = math.Pow(3, 4)
	const c33 = len("hello") // this works since len is a function which is available at compile time also //1. defined string is unnamed constants  // 2. len function uses a constant as an input, and the input string was constant only // 3. hence the output of the function can be determined at compiler time only
	fmt.Println(c33)

	// data typed constansts are strict
	const aa int = 4
	// const bb float32 = 2.2 * aa

	// data typeless are flexible
	const aaa = 4
	const bbb = 2.2 * aaa // this doesnt change the type of aaa // but when initizing bbb // go compiler itself call the float // so it looks like bbb = 2.2 * float64(aaa)
	fmt.Printf("a : %T   b: %T \n", aaa, bbb)

	/*
		iota // this is just to keep constant/enum declaration clean (its some fun feature to remove a lil bit boilerplate)

		value of iota starts with 0 and is incremented every time its called
	*/

	const (
		NORTH1 int = 0
		EAST1  int = 1
		SOUTH1 int = 2
		WEST1  int = 3
	)
	fmt.Printf("N:%d E:%d S:%d W:%d\n", NORTH1, EAST1, SOUTH1, WEST1)

	const (
		NORTH22 = iota
		// EAST2 int // wont work // iota works on untyped constants
		EAST22 = iota
	)
	fmt.Printf("N:%d E:%d \n", NORTH22, EAST22)

	const (
		NORTH2 = iota
		EAST2  // this is the feature of constant // omitting
		SOUTH2
		WEST2
	)
	fmt.Printf("N:%d E:%d S:%d W:%d\n", NORTH2, EAST2, SOUTH2, WEST2)

	// whenever next "const" keywork appears // iota RESETS to 0
	const (
		NORTH3 = iota + 1
		EAST3
		SOUTH3
		WEST3
	)
	fmt.Printf("N:%d E:%d S:%d W:%d\n", NORTH3, EAST3, SOUTH3, WEST3)

	const (
		NORTH32 = iota * 2
		EAST32
		SOUTH32
		WEST32
	)
	fmt.Printf("N:%d E:%d S:%d W:%d\n", NORTH32, EAST32, SOUTH32, WEST32)

	const (
		NORTH33 = 1 << iota
		EAST33
		SOUTH33
		WEST33
	)
	fmt.Printf("N:%b E:%b S:%b W:%b\n", NORTH33, EAST33, SOUTH33, WEST33)

	// WHEN NOT TO USE IOTA
	// we basically use iota to generate some code for us // help in initializing enum/constant values
	// if the consntnat/ enum values will be persisted in db
	// e.g. compiled code after using iota NORTH 0, EAST : 1 SOUTH : 2 WEST: 3
	// db row : {City(name: "bengaluru", dir : 2)}
	// BENGALURU in SOUTH when db row read in code/app
	// lessay and later you added one more constant NORTH_EAST b/w NORTH and EAST //
	// e.g. compiled code after using iota NORTH 0, NORTH_EAST : 1, EAST : 2 SOUTH : 3 WEST: 4
	// db row : {City(name: "bengaluru", dir : 2)}
	// BENGALURU in EAST when db row read in code/app
	// ordering and number value will now change for all enums // deserialization of all values in db will now break :O  Nooooooooooo!
	// we actually wanted the mapping like NORTH 0, EAST : 1 SOUTH : 2 WEST: 3 NORTH_EAST : 5
	// and backward compatible consistent values  BENGALURU in SOUTH
}
