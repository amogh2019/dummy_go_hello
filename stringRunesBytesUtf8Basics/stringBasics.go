package main

import (
	"fmt"
	"strings"
	"unicode/utf8"
)

func main() {

	// backticks doesnt do string processing at all // does not need to escape the double quotes and what not
	fmt.Println(`hello! "YOLO" \n oye luckky luckky oye!`)

	// About char in Go
	// there is no char datatype in Go
	// characters are bacisally represented by their unicode value(unicode values representating that character in that unicode-charater table)
	// hence a char may take either 1 byte or 4 byte
	// so a character is either represented as byte or rune in Go
	// byte === alias for uint8
	// rune === alias for int32
	// So, strings in Go ===[actual] sequence of bytes ~=[abstract] sequence of runes  != sequence of char(since there is no char)

	s1 := "abcd"
	fmt.Println(s1, len(s1), utf8.RuneCountInString(s1))

	// len of string // returns the no. of bytes // not the number of chars/runes
	s1 = "abकखगघ"
	fmt.Println(s1, len(s1), utf8.RuneCountInString(s1)) // here length is not 4 // each char's unicode represenation take 3 bytes == total 12 bytes

	// // iterating strings // 1 // byte by byte
	// fmt.Println(strings.Repeat("#", 30))
	// for i := 0; i < len(s1); i++ {
	// 	fmt.Printf("%c ", s1[i]) // since this is printing byte by byte // its not considering that the charater 'क' actually takes three bytes
	// } //TODO not sure why // uncommenting this for loop breaks things

	// iterating strings // 2 // rune by rune
	fmt.Println(strings.Repeat("#", 30))
	for i := 0; i < len(s1); {
		prefixRune, prefixRuneSize := utf8.DecodeRuneInString(s1[i:])
		fmt.Printf("%c ", prefixRune)
		i += prefixRuneSize
	}

	// iterating strings // 3 // using range to iterate
	fmt.Println("\n" + strings.Repeat("#", 30))
	for _, rune := range s1 {
		fmt.Printf("%c ", rune)
	}

}
