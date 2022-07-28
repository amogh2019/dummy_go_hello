package main

import (
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"os"
)

func main() {

	fileName := "go.mod"
	// READING with io // into buffer // not expecting EOF end of file character in buffer
	file, err := os.Open(fileName)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	dataSlice := make([]byte, 20)
	bytesRead, err := io.ReadFull(file, dataSlice)
	if err != nil {
		log.Fatal("buffer size is greater than file size, we were expecting to completely fill the buffer with the file data ", err)
	}
	fmt.Println("bytes read when buffer size to read is fixed : ", bytesRead, string(dataSlice))

	fullFileInBuffer, err := io.ReadAll(file)
	fmt.Println("complete file in buffer created by io package only,  continued from last read point in file\n", string(fullFileInBuffer))

	sameFile, _ := os.Open(fileName)
	defer sameFile.Close()
	fullFileInBuffer, err = io.ReadAll(sameFile)
	fmt.Println("complete file in buffer created by io package only, from starting point\n", string(fullFileInBuffer))

	// READING with ioutils + manual File management
	sameFile, _ = os.Open(fileName)
	defer sameFile.Close()
	fullFileInBuffer, err = ioutil.ReadAll(sameFile)
	fmt.Println("complete file in buffer created by io package only, from starting point\n", string(fullFileInBuffer))


	// READING with ioutils + no file management
}
