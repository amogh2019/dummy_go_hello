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
		log.Fatal(err) // === print + exit // this also prints a timestamp
	}

	//CLEAR FILE
	err = os.Truncate("abc.txt", 0) // shorten file to a size of bytes specified // here 0, means clear the file
	if err != nil {
		log.Fatal("error while clearing file", err)
	}

	//OPEN FILE
	file1, err := os.Open("abc.txt")
	if err != nil {
		log.Fatal("cannot open", err)
	}
	file1.Close()
	file1, _ = os.OpenFile("abc.txt", os.O_APPEND|os.O_CREATE, 0644) // either open in append mode if exists // or create the file if doesnt exist
	file1.Close()

	// FILE INFO
	fileInfo, err := os.Stat("abc.txt")
	fmt.Println(fileInfo)
	fmt.Println(fileInfo.Name())
	fmt.Println(fileInfo.Size())
	fmt.Println(fileInfo.IsDir())
	fmt.Println(fileInfo.ModTime())

	// fileInfo, err = os.Stat("asdfasfabc.txt")
	// if err != nil {
	// 	if os.IsNotExist(err) { // reading the error and taking decision based on error value
	// 		log.Fatal("this file does not exist ", err)
	// 	}
	// }

	// RENAMING
	oldName := "abc.txt"
	newName := "abc1.txt"
	err = os.Rename(oldName, newName)

	if err != nil {
		log.Fatal("error in renaming ", err)
	}
	fmt.Println("renamed")

	// REMOVE
	err = os.Remove(newName)
	if err != nil {
		log.Fatal("error in remoing ", err)
	}
	fmt.Println("removed")

}
