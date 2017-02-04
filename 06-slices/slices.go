// https://golang.org/ref/spec#Slice_types
// https://golang.org/doc/effective_go.html#allocation_make

// the underlaying type for a slice is an array. Think about slices
// as an abstraction on top of arrays. The nice thing with slices is
// that they don't need to have a specifid capacity its length
// can dynamically be increased

package main

import "fmt"

func main() {
	// create a slice using composite literals
	x := []string{"everything", "is", "data"}

	// give me a slice of items from index 0 to 1
	everything := x[:1]
	fmt.Println(everything)
	// give me a slice of items from index 2 to end
	data := x[2:]
	fmt.Println(data)

	// we can also use the make function to allocate a slice
	y := make([]string, 10)
	fmt.Println(len(y))

	// append slices to a slice
	// give me index 1 but leave out index 2
	y = append(x[1:2])

	// slices can be copied
	clone := make([]string, len(x))
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
