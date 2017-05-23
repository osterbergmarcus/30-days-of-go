/*
* List is a linear data structure with a starting point and ending point, it
* contains data in a ordered sequence
* this is handy for dealing with elements at the begining or end of the collection
* we will work with slices here since they are flexible and extensible in comparsion to Arrays
 */

package list

// List - this will be our structure
type List struct {
	memory []interface{}
	length int
}

// Push - adds a value to the end of the List
func (list *List) Push(value interface{}) {
	/*
	* alocates new memory for our list with increased capacity
	* copy over existing values and add new values to list
	* the built in function append will do the hard work
	 */
	list.memory = append(list.memory, value)

	// increase the length of the List
	list.length++
}

// Pop - pop the last value from the List
func (list *List) Pop() interface{} {
	// if the memory is empty then terminate
	if list.length == 0 {
		return nil
	}

	// decrease length and shrink memory by one, also return popped value
	lastIndex := list.length - 1
	lastIndexValue := list.memory[lastIndex]
	list.memory = list.memory[:lastIndex]
	list.length = lastIndex

	return lastIndexValue
}

// Unshift - to add a value to the begining of the List, we need to move all current values to the side
func (list *List) Unshift(value interface{}) {
	newMemory := []interface{}{value}

	// using the operator  ... lets you spread out multiple values
	list.memory = append(newMemory, list.memory...)

	list.length++
}

// Shift - removes the first value from the List and moves remaining values back one step
func (list *List) Shift() interface{} {
	// if the memory is empty then terminate
	if list.length == 0 {
		return nil
	}

	// decrease length and shrink memory by one, also return shifted value
	firstIndexValue := list.memory[0]
	list.memory = list.memory[1:]

	list.length--

	return firstIndexValue
}

// Get - get returns a specific value based on its index
func (list *List) Get(index int) interface{} {
	return list.memory[index]
}
