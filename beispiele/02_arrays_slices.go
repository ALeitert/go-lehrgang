package foundations

import "fmt"

// Array vs Slice:
//  * Array has fixed size.
//  * Slice references a part of some array.

func arrays() {
	// Array with zero values.
	var arr [4]int

	// With initial values.
	b := [4]int{0, 1, 2, 3}

	// If not all values are declared, they become zero.
	// Bad habit. Do not do that!
	c := [4]int{0, 1, 2 /*, 0 */}

	// Set size at compile time:
	d := [...]int{0, 1, 2, 3}
}

func slices() {
	// Slices are fat pointers to underlying arrays.
	// Go will manage arrays for us.
	//
	//	slice: [ * | len: 3 | cap: 5 ]
	//           |
	//           +---------+
	//                     ▼
	//  array: [   |   |   | 0 | 1 | 2 | _ | _ ]
	//
	// len: length of the slice, values we can access
	// cap: capacity, available space in underlying array (after start of slice)

	// Create a slice.
	a := []int{0, 1, 2, 3}
	fmt.Println(a[1]) // 1
	a[1] = 5
	fmt.Println(a[1]) // 5

	// We can also use `make()`.
	b := make([]int, 5 /* len */)
	c := make([]int, 0 /* len */, 10 /* cap */)

	//
	// Slice operation:
	//	slice[ start (inclusive) : end (exclusive) ]

	bigArr := []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}

	sub1 := bigArr[:4]  // { 0, 1, 2, 3 }
	sub2 := bigArr[3:5] // { 3, 4 }
	sub3 := bigArr[7:]  // { 7, 8, 9 }

	// Important: They share the underlying array.
	sub1[3] = 42
	fmt.Println(sub2[0]) // 42

	// len() and cap()
	fmt.Println(len(bigArr), cap(bigArr)) // 10  10
	fmt.Println(len(sub1), cap(sub1))     //  4  10
	fmt.Println(len(sub2), cap(sub2))     //  2   7
	fmt.Println(len(sub3), cap(sub3))     //  3   3

	//
	// Empty vs `nil` slice

	var (
		eSl = []int{}
		nSl []int
	)

	fmt.Println(eSl == nil, len(eSl), cap(eSl)) // false, 0, 0
	fmt.Println(nSl == nil, len(nSl), cap(nSl)) //  true, 0, 0
}

func growingSlices() {
	//
	// Add with append()

	slice := []int{0}
	slice = append(slice, 1)
	fmt.Println(slice) // [ 0, 1 ]

	// Important: `append` returns a new slice. It does not modify the given.
	newSlice := append(slice, 2)
	fmt.Println(slice)    // [ 0, 1 ]
	fmt.Println(newSlice) // [ 0, 1, 2 ]

	//
	// Capacity grows (doubles) as needed.

	slice = append(slice, 2)
	fmt.Println(len(slice), cap(slice)) // 3  4

	slice = append(slice, 3)
	fmt.Println(len(slice), cap(slice)) // 4  4

	slice = append(slice, 4)
	fmt.Println(len(slice), cap(slice)) // 5  8

	//
	// Queue

	q := []int{}

	// enqueue
	q = append(q, 1)

	// dequeue
	front := q[0]
	q = q[1:]

	//
	// Stack

	stack := []int{}

	// push
	stack = append(stack, 1)

	// pop
	top := stack[len(stack)-1]
	stack = stack[:len(stack)-1]
}
