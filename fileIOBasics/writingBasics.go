package main

import (
	"bufio"
	"fmt"
	"io/ioutil"
	"log"
	"os"
)

func main() {

	// STEPS to write to a file
	// 1. open file
	// 	    a. decide on what will be the mode of opening // create new // append existing file mode // clear existing file
	// 2. write byte slice on the file
	// 3. close the file

	// write by os package // lightweight
	file, err := os.OpenFile("aa.txt", os.O_CREATE|os.O_APPEND|os.O_WRONLY, 0644)
	// file, err := os.OpenFile("aa.txt", os.O_CREATE|os.O_TRUNC|os.O_WRONLY, 0644) // this will clear the exiting file // WRTIE-ONLY mode?
	if err != nil {
		fmt.Println("error in opening file ", err)
	}
	defer file.Close()

	data := []byte("This is sparta\n") // string to byte slice
	bytesWritten, err := file.Write(data)
	if err != nil {
		fmt.Println("error in writing data to file", file, " ", err)
	}
	fmt.Println("bytes written is ", bytesWritten)

	// write by ioutils
	// increase the size of executable
	// has more utility functions
	// doesnt require explicit file closing
	// this handles // creating // opening // writingslice // closing automatically
	err = ioutil.WriteFile("bb.txt", data, 0664)
	if err != nil {
		fmt.Println("error in writing data to file", file, " ", err)
	}

	// write by buffered writer
	// buffered writer maintains a size that holds the data to be written
	// when the batch is full, it writes the data on the file and reset
	// we can also manually flush and reset
	file2, err := os.OpenFile("cc.txt", os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		log.Fatal(err)
	}
	defer file2.Close()

	bufferedWriter := bufio.NewWriterSize(file2, 5)
	fmt.Println("buffer size", bufferedWriter.Available())

	tempstr := "abcdefghijk"
	for _, singleChar := range tempstr {
		fmt.Println([]byte{byte(singleChar)})
		bufferedWriter.Write([]byte{byte(singleChar)})

		fmt.Println("buff", bufferedWriter.Available())
		fmt.Println("buff", bufferedWriter.Buffered())
	}
	if bufferedWriter.Available() > 0 {
		bufferedWriter.Flush()
	}

}
