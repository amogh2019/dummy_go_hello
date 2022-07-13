package main

import (
	"fmt"
	"time"
)

func main() {

	// different from other languagues
	// not need to write breaks
	// each cases breaks by default
	// switch case here is like simplified if else only // its translated to that if else only
	// type to compare with the condition must be same
	// default is true that is compared for each case

	name := "go"
	switch name {
	case "c":
		fmt.Println("selected language is c")
	case "go":
		fmt.Println("selected language is go")
	default:
		fmt.Println("selected language is dunno")
	}

	currentHour := time.Now().Hour()
	switch true {
	case currentHour < 5:
		fmt.Println("hours is less than 12")
	case currentHour < 12:
		fmt.Println("hours is less than 12")
	default:
		fmt.Println("hours is greater than 12")
	}

	// default value to check is true // means the case to be compared should be boolean type
	// works without default case,
	i := 22
	switch {
	case i%2 == 0:
		fmt.Println("i", i, "is even")
	case i%2 != 0:
		fmt.Println("i", i, "is odd")
	}

}
