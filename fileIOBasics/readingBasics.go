package main

import (
	"bufio"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"os"
	"strings"
)

func main() {

	fileName := "go.mod"
	// READING
	// with io // into buffer // not expecting EOF end of file character in buffer
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

	fullFileInBuffer, _ := io.ReadAll(file)
	fmt.Println("complete file in buffer created by io package only,  continued from last read point in file\n", string(fullFileInBuffer))

	sameFile, _ := os.Open(fileName)
	defer sameFile.Close()
	fullFileInBuffer, _ = io.ReadAll(sameFile)
	fmt.Println("complete file in buffer created by io package only, from starting point\n", string(fullFileInBuffer))

	// READING
	// with ioutils + manual File management
	sameFile, _ = os.Open(fileName)
	defer sameFile.Close()
	fullFileInBuffer, _ = ioutil.ReadAll(sameFile)
	fmt.Println("complete file in buffer created by io package only, from starting point\n", string(fullFileInBuffer))

	// READING
	// with ioutils + no file management
	fileData, err := ioutil.ReadFile(fileName)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(string(fileData))
	fmt.Println(strings.Repeat("###", 20))

	// BUFFERED READING
	// open file // attach a buffered reader to the file
	// read buffer by buffer till scanner.err != nil   // here if EOF is reached // scanner's err becomes nil and overall scan is false // to let you know that cannot read anymoh
	f1, err := os.Open(fileName)
	if err != nil {
		log.Fatal(err)
	}
	defer f1.Close()

	buffReader := bufio.NewScanner(f1) // by default buffer breaks the data by '\n' character // you can let the buffer be broken as per your own delimiter
	// buffReader.Split(bufio.ScanRunes) // delimiter == end of each rune // char by char reading
	// buffReader.Split(bufio.ScanWords) // delimiter == end of each word // word by word reading

	i := 0
	for ; buffReader.Scan() != false; i++ { // this scans for new line // returns true if line is read // returns false if EOF or err
		fmt.Println(i, buffReader.Text())
		// fmt.Println(i, string(buffReader.Bytes()))
	}
	if err := buffReader.Err(); err == nil {
		fmt.Println("File", fileName, "is read completely, buffered reader scan returned false, and buffered reader error is nil, means it reached EOF")
	} else if err != nil {
		log.Fatal("some error in buffered reading", err)
	}

	fmt.Println(strings.Repeat("###", 20))
}
