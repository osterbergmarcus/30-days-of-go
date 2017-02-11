// https://golang.org/ref/spec#Slice_types
// https://golang.org/doc/effective_go.html#allocation_make

// the underlaying type for a slice is an array. Think about slices
// as an abstraction on top of arrays. The nice thing with slices is
// that they don't need to have a specifid capacity its length
// can dynamically be increased

package main

import "fmt"

func main() {
	type User struct {
		name string
		role string
		id   int
	}

	// create a slice using composite literals
	x := []User{
		{"User1", "admin", 1},
		{"User2", "basic", 2},
		{"User3", "basic", 3},
	}
	fmt.Println(x)

	// give me a slice of items from index 0 to 1
	someUser := x[:1]
	fmt.Println(someUser)

	// give me a slice of items from index 2 to end
	anotherUser := x[2:]
	fmt.Println(anotherUser)

	// we can also use the make function to allocate a slice
	y := make([]string, 10)
	fmt.Println(len(y))

	// append slices to a slice
	// give me index 1 but leave out index 2
	someUser = append(x[1:2])

	// slices can be copied
	clone := make([]User, len(x))
	copy(clone, x)
	fmt.Println("Cloned slice:", clone)

	// loop over a slice
	loop := make([]int, 10)
	for index := range loop {
		loop[index] = index * 2
	}
	fmt.Println("index x 2:", loop, len(loop))

	// slices can be multi-dimensionals
	// slice within a slice
	multiDimension := make([][]string, 1)
	inception := []string{"inception"}
	multiDimension[0] = inception
	fmt.Println("Multi-dimension slice:", multiDimension)

	// slices are homogeneous
	// its items can only be of the same type
	// we can use an empty interface to create a hetrogeneous slice
	// a slice where its items are of different types
	differentKinds := []interface{}{"foo", 0, true}
	fmt.Println("Different types:", differentKinds)
}
