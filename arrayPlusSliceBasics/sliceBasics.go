package main

import (
	"fmt"
	"strings"
	"unsafe"
)

func main() {

	// slices are  vectors // they are dynamic and len is known at runtime time only
	// slics are  like List<> pointers // each slics, since if not initialized, points to 'nil' //  else would be to some memory location

	// simple declaration nil init
	var numbers []int             // variable // named numbers // of type  sliceSinceSizeIsUnknownOfTypeInt
	fmt.Printf("%#v \n", numbers) // empty slice with zero capacity
	fmt.Println(len(numbers))     // 0
	// numbers[1] = 13            // cannot assign since not initialized // numbers is pointing to nil // capactiy is zero for now

	// simple init with zero
	var s1 = make([]int, 5)
	fmt.Printf("%#v \n", s1)

	// simple init with values
	var s2 = []int{1, 2, 4, 22, 3}
	fmt.Printf("%#v \n", s2)
	s2[0] = 1000
	fmt.Printf("%#v \n", s2)
	// s2[10] = 1022             // increase length first

	// SLICE COMPARISON

	var s11 []int            // just declared // slice uninit // equals to nil
	fmt.Println(s11 == nil)  // true
	fmt.Printf("%#v\n", s11) // []int(nil)
	fmt.Println(len(s11))    // 0 since nil

	s12 := []int{}          // init // equals to empty
	fmt.Println(s12 == nil) // false
	fmt.Println(len(s12))   // 0 since empty

	// slices can only be compared to nil using == operator
	// else write a function
	a11 := []int{1, 2, 3}
	a12 := []int{1, 2, 3}
	a13 := []int{1, 3, 2}
	a14 := []int{1, 2, 3, 4}
	a15 := []int{1, 2}
	var a16 []int
	fmt.Println(compareSlices(a11, a12)) // true
	fmt.Println(compareSlices(a11, a13)) // false
	fmt.Println(compareSlices(a11, a14)) // false
	fmt.Println(compareSlices(a11, a15)) // false
	fmt.Println(compareSlices(a11, a16)) // false
	// other way around
	fmt.Println(compareSlices(a12, a11)) // true
	fmt.Println(compareSlices(a13, a11)) // false
	fmt.Println(compareSlices(a14, a11)) // false
	fmt.Println(compareSlices(a15, a11)) // false
	fmt.Println(compareSlices(a16, a11)) // false

	// append to slice
	// append returns a new slice
	// append takes dest and then the list of elements
	a33 := []int{1, 2, 3}
	a33 = append(a33, 11, 33, 44)
	fmt.Println(a33)
	var a331 = make([]int, 2) // [0, 0]
	a331 = append(a331, a33...)
	fmt.Println(a331) // [0, 0, 1,2, .....]

	// copy to destination from source
	// this doesnt not clone // only copies on destination till the size of destination
	a44 := []int{1, 2, 3, 4}
	a441 := make([]int, 2)                //[ 0 , 0 ]
	var countCopied int = copy(a441, a44) // only copies two element since destination size is 2
	fmt.Println(a44, a441, countCopied)

	// slice expressions
	// array slices and string are slicable
	ss11 := []int{1, 2, 3, 4}
	fmt.Println(ss11)
	fmt.Println(ss11[1:2]) // starting index // excluding this last index
	// fmt.Println(ss11[1:100]) // error
	fmt.Println(ss11[1:]) // == ss11[1:len(ss11)]
	fmt.Println(ss11[:2]) // == ss11[0:2] // include 0th index // exclude 2nd index
	fmt.Println(ss11[:])  // == ss11[0:len(ss11)]
	ss12 := ss11[1:2]     // this does not return a new slice // instead slice uses the same backing array of the orignal // hence changing element here change the original element as well
	fmt.Println(ss12)
	fmt.Println(len(ss12))
	fmt.Println(cap(ss12))
	fmt.Println(append(ss11[:3], 1000))             // slice returns new // then append also returns  a new slice
	fmt.Println(append([]int{22, 33}, ss11[:3]...)) // slice returns new // then append also returns  a new slice //  shared backing array wala issue complex // when elems added it creates a backing array of cap twice of previous , so might not create new backing array every time

	ss13 := ss11[0:]
	fmt.Println("old val", ss11[0])
	fmt.Println("changing the values of the sliced expression")
	ss13[0]++
	fmt.Println("new val", ss11[0])
	fmt.Println("original changed since backing array of slice is same")
	fmt.Printf("slice headers different %p %p  but backing array shared\n", &ss11, &ss13)

	// slice = sliceHeader+  a backing array (hidden)
	// slice header =  backing array pointer + len() backing array orignal length + cap()  number of elements in the underlying array, counting from the first element in the slice.   == slice metadata
	// nil slice has no backing array

	// printing the memory occupied
	arr1 := [5]int{1, 2, 3, 4, 5}                                   // array of size 5
	sl1 := []int{1, 2, 3, 4, 5}                                     // slice of size 5
	fmt.Printf("array : size in bytes : %d\n", unsafe.Sizeof(arr1)) // 8 bytes * 5 = 40
	fmt.Printf("slice : size in bytes : %d\n", unsafe.Sizeof(sl1))  // 8 bytes * 3 = 24  (pointer to backing arary, length, cap) // slice === slice header(slice metadat)(the ds) // backing array size is not includede

	//  append()'s behaviour of creating a new backing array
	fmt.Println(strings.Repeat("#####", 4))
	ss22 := []string{"A", "B", "C", "D"}
	fmt.Println("orignal arr", ss22)
	ss23 := append(ss22[:2], "X", "Y")
	fmt.Println("append changed backing array of original", ss22, ss23, "not creating a new backing array since backing array cap can handle this load")
	fmt.Printf("original slice %#v, len %d, cap %d\n", ss22, len(ss22), cap(ss22))
	fmt.Printf("appended slice %#v, len %d, cap %d\n", ss23, len(ss23), cap(ss23))

	fmt.Println(strings.Repeat("#####", 4))
	ss22 = []string{"A", "B", "C", "D"}
	fmt.Println("orignal arr", ss22)
	ss23 = append(ss22[:2], "X", "Y", "Z", "P")
	fmt.Println("append DID NOT change backing array of original", ss22, ss23, "not creating a new backing array since backing array cap can handle this load")
	fmt.Printf("original slice %#v, len %d, cap %d\n", ss22, len(ss22), cap(ss22))
	fmt.Printf("appended slice %#v, len %d, cap %d\n", ss23, len(ss23), cap(ss23))

	fmt.Println(strings.Repeat("#####", 4))
	orig := make([]string, 10)
	ss22 = orig[:5]
	fmt.Println("orignal arr", ss22)
	ss23 = append(ss22[:2], "X", "Y", "Z", "P")
	fmt.Println("append changed backing array of original", ss22, ss23, "not creating a new backing array since backing array cap can handle this load")
	fmt.Printf("original slice %#v, len %d, cap %d\n", ss22, len(ss22), cap(ss22))
	fmt.Printf("appended slice %#v, len %d, cap %d\n", ss23, len(ss23), cap(ss23))

	// slicesExpression taken from backing array not slice
	s44 := []int{1, 2, 3, 4, 5}
	s43 := s44[:2] // has 1,2 // idx 0, 1
	// fmt.Println(s43[3]) //index out of range [3] with length 2
	// but we can slice from backing arary
	fmt.Println(s43[3:5])
}

func compareSlices(a []int, b []int) bool {
	if a == nil && b == nil {
		return true
	}

	if a == nil || b == nil {
		return false
	}

	// compare their lengths and values in order
	if len(a) != len(b) {
		return false
	}

	for idx, val := range a {
		if b[idx] != val {
			return false
		}
	}
	return true
}
