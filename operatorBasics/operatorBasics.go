package main

import "fmt"

func main() {

	a, b := 4, 2

	// follows operator precedence  // parenthses then disivison then mul
	r := (a + b) / (a - b) * 2
	fmt.Println(r) // 6

	// mod opreator
	r = 9 % a
	fmt.Println(r) // 1

	// assignment operator
	a, b = 2, 3
	a += b
	fmt.Println(a) // 5
	b -= 2
	fmt.Println(b) // 1
	b *= 10
	fmt.Println(b) // 10
	b /= 5
	fmt.Println(b) // 2
	a %= 3
	fmt.Println(a) // 2

	// increment/decrement is not an operator in Go, its a statement
	a = 10
	a += 1
	a++
	fmt.Println(a) // 12
	// fmt.Println(a++) // illegal

	// conditional opreator + logical operators
	// a = 12 b = 2
	fmt.Println("a:", a, "b:", b) // true
	fmt.Println(a > 1 && b > 1)   // true
	fmt.Println(a < 1 && b > 1)   // false
	fmt.Println(a < 1 || b > 1)   // true

	// two more operators
	// oprator for pointers
	// opreator for channels






	
}
