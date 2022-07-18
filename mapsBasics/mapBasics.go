package main

import (
	"fmt"
	"strings"
)

func main() {

	// maps // a collection type
	// all keys must have same type // all values must have same type // type of key and value can be diff
	// keys can only be those type which are fully comparable by == or != // e.g. string int array // slices are not fully comparable, can only be compared to nil
	// maps are not fully comparable  // can only be compared to nil
	// unordered keys // maps are unordered data structure in go

	// nil map
	var m11 map[string]string
	fmt.Printf("%#v", m11)
	fmt.Printf("values for %q in map is %q\n", "hey", m11["hey"])

	// int maps // unfound keys are values zero not nil in int maps // in string values its empty string
	var m12 map[string]int
	fmt.Printf("values for %q in map is %d\n", "hey", m12["hey"])

	// only comparable types as keys
	// var m13 map[[]int]int // slices cannot be keys // since slices are not fully comprable
	var m14 map[[4]int]int
	// _ = m13
	_ = m14

	fmt.Println(strings.Repeat("#", 30))

	// maps needs to be initialized after declaration
	// m11["key"] = "value" // will break untill m11 is initialized // m11 = map[string]string{}

	// init 1 // empty map
	age := make(map[string]int)
	fmt.Println(age)
	fmt.Println(age["saitama"])

	// init 2 // without make // empty map
	age1 := map[string]int{}
	fmt.Println(age1)
	fmt.Println(age1["saitama"])

	// i1nit 3 // without make // non empty map
	age11 := map[string]int{
		// 111 : 11000, // incorrect key // all key types must be same
		"saitama": 100000000, // we give extra comma in go at last line // in multi line initialization
	}
	fmt.Println(age11)
	fmt.Println(age11["saitama"])

	// map // contains method
	// since we know, we if check for a key in map, and it does not exisit, that map returns 0 or the empty value representation
	// to check if the values was actually zero or its not present, we can harness the extra value that it returns
	saitamaAge, isPresent := age11["saitama"]
	if isPresent {
		fmt.Println("saitama is present and its age is", saitamaAge)
	} else {
		fmt.Println("saitama not present")
	}
	codex, isPresent := age11["codex"] // am able to use isPresent again, since short declaration requires atleast one new variable declaration, and that is codex in our case
	if isPresent {
		fmt.Println("codex is present and its age is", codex)
	} else {
		fmt.Println("codex not present")
	}

	// iterating map

	fmt.Println(strings.Repeat("#", 30))
	age11["codex"] = 24
	for k, v := range age11 {
		fmt.Println(k, ":", v)
	}

	fmt.Println(strings.Repeat("#", 30))

	// deleting key from map
	age11["codex11"] = 24
	fmt.Printf("%#v\n", age11)
	delete(age11, "codex11") // no error if key is not present // can avoid checking key presence
	fmt.Printf("%#v\n", age11)

	// map  comparison
	// cannot compare m1 == m2
	// can only do m1 == nil

	// map cloning 1 // incorrect cloning
	fmt.Println(strings.Repeat("#", 30))
	m111 := map[string]int{"saitama": 1000}
	m112 := m111 /// they share the same ds // both pointers point to same backing structure
	m112["codex"] = 11
	fmt.Println(m111, m112)

	// map cloning 1 // deep cloning // manual loop
	fmt.Println(strings.Repeat("#", 30))
	m113 := make(map[string]int)
	for k, v := range m111 {
		m113[k] = v
	}
	m113["newcodex"] = 3
	fmt.Println(m111, m113)
}
