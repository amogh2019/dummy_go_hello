package main

import (
	"fmt"
	"log"
	"os"
)

func main() {

	// Packages that we may need to work with files
	// os // io // ioutils // bufio

	// use of os write file
	var newFile1 *os.File
	fmt.Printf("%T\n", newFile1)

	// CREATE // new if not existing // truncate(clear) if it already exists
	var err error // variable of type error
	newFile1, err = os.Create("abc.txt")
	// newFile1, err = os.Create("/asdf/abc.txt") // this will give no directory error
	if err != nil {
		// fmt.Println(err)
		// os.Exit(1)
		log.Fatal(err) // === print + exit
	}
}
