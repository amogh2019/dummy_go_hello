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
	fmt.Println("\n" + strings.Repeat("#", 30))

	// slicing a string
	// since we know /// string === []bytes
	// 1. so when we slice string , we get the slice of bytes // this would break since a char can be three byte large and we tried to get two btyes // but this is efficient since sliceing here won't create a new backing array
	// 2. we convert string to rune-slice // i.e. string == []bytes => []runes // this would help us in easier slicing, since each charater is one rune, nothing breaks now // but this is memory expensice // creating a new rune slice from byte slice creates new backing array obvio
	fmt.Println(s1[0:6]) // abक�  // a(1byte) b(1byte) क(3bytes) �(ख's first byte, unmapped unicode-codepoint to errorcharacter, ख takes three bytes)
	runeSliceOfS1 := []rune(s1)
	fmt.Println(runeSliceOfS1[0:4])         // prints integers // since its slice of rune(int32)
	fmt.Println(string(runeSliceOfS1[0:4])) // abकख
	fmt.Println("\n" + strings.Repeat("#", 30))

	// strings library functions

	p := fmt.Println
	//Contains //  it returns true whether a substr is within a string
	p(strings.Contains("go golang gopher", "golang"))

	//ContainsAny //  // it returns true whether any Rune / Unicode code points are within our string, and false otherwise.
	p(strings.ContainsAny("go golang gopher", "xyp"))

	// it reports whether a rune is within a string.
	p(strings.ContainsRune("go golang gopher", 'p')) // true
	p(strings.ContainsRune("go golang gopher", 'x')) // false

	// it returns the number of instances of a substring in a string
	p(strings.Count("go golang gopher", "go")) // 3

	// if the substr is an empty string Count() returns 1 + the number of runes in the string
	p(strings.Count("five", "")) // 5

	// upper lower
	p(strings.ToLower("FivE fouR")) // five four
	p(strings.ToUpper("FivE fouR")) // FIVE FOUR

	// comparing strings (case matters)
	// p("go" == "go") // -> true
	p("Go" == "go") // -> false

	// string comparison // case insensetive // expensive since toLower converts the string into lower, rune by rune
	p(strings.ToLower("Go") == strings.ToLower("go")) // even ide gives warning

	// string comparison // case insensetive // without converting to lower
	p(strings.EqualFold("GO", "go")) // true

	// replace // 2 occurences// all occurense // replceAll mehtod
	p(strings.Replace("192.168.0.1", ".", ":", 2))  // it replaces the first 2 occurrences
	p(strings.Replace("192.168.0.1", ".", ":", -1)) // it replaces the first 2 occurrences
	p(strings.ReplaceAll("192.168.0.1", ".", ":"))  // it replaces the first 2 occurrences

	// split
	s := strings.Split("a,b,c", ",")
	fmt.Printf("%T\n", s)                                                // -> []string slice of strings
	fmt.Printf("strings.Split():%#v\n", s)                               // => strings.Split():[]string{"a", "b", "c"}
	fmt.Printf("strings.Split():%#v\n", strings.Split("Go for Go!", "")) // -> []string{"G", "o", " ", "f", "o", "r", " ", "G", "o", "!"}

	//Join
	// The separator string is placed between elements in the resulting string.
	s = []string{"I", "learn", "Golang"}
	j := strings.Join(s, "-")
	fmt.Printf("%T\n", j) // -> string
	p(j)                  // -> I-learn-Golang

	// splitting a string by whitespaces and newlines.
	myStr := "Orange Green \n Blue Yellow"
	fields := strings.Fields(myStr) // it returns a slice of strings
	fmt.Printf("%T\n", fields)      // -> []string
	fmt.Printf("%#v\n", fields)     // -> []string{"Orange", "Green", "Blue", "Yellow"}

	// TrimSpace() removes leading and trailing whitespaces and tabs.
	s1 = strings.TrimSpace("\t Goodbye Windows, Welcome Linux!\n ")
	fmt.Printf("%q\n", s1) // "Goodbye Windows, Welcome Linux!"

	// To remove other leading and trailing characters, use Trim()
	s2 := strings.Trim("...Hello, Gophers!!!?", ".!?")
	fmt.Printf("%q\n", s2) // "Hello, Gophers"
}
