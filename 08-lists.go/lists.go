// if you are familiar with Arrays in JavaScript or Lists in any other languages
// then you might know that a List contains your data in a ordered sequence
// you can dynamically add and remove items

// List is a linear data strucutre with a starting point and ending point

package list

// List - this will be our structure
type List struct {
	memory []interface{}
	length int
}

// Push - we need a method to add a value to the end of our List
func (list *List) Push(value interface{}) {
	// since our List has current capacity of 0
	// we will need to re allocate the memory
	currentCapacity := cap(list.memory)
	newCapacity := currentCapacity + 1
	currentLength := list.length
	newLength := currentLength + 1
	newMemory := make([]interface{}, newLength, newCapacity)

	// clone old memory to new allocated memory
	copy(newMemory, list.memory)
	list.memory = newMemory

	// add new value to the last index of our List
	// and increase the length of our List
	list.memory[currentLength] = value
	list.length = newLength
}

// Get - Get will give us a specific value based on its index
func (list *List) Get(index int) interface{} {
	return list.memory[index]
}
