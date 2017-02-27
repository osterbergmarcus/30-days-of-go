/*
* if you are familiar with Arrays in JavaScript or Lists in any other languages
* then you might know that a List contains your data in a ordered sequence
* you can dynamically add and remove values
* Lists are very useful when you are dealing with elements at the begining or end of your collection
*
* List is a linear data structure with a starting point and ending point
* we will work with slices here since they are flexible and extensible in comparsion to Arrays
 */

package list

// List - this will be our structure
type List struct {
	memory []interface{}
	length int
}

// Push - we need a method to add a value to the end of our List
func (list *List) Push(value interface{}) {
	/*
	* we need to realocate new memory for our list with increased capacity
	* copy over existing values and add new values to list
	* the built in function append will do the hard work for us
	 */
	list.memory = append(list.memory, value)

	// increase the length of our List
	list.length++
}

// Pop - we need a method to pop the last value from the List
func (list *List) Pop() interface{} {
	// if the memory is empty then terminate function
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

// Unshift - to add an value at the begning of our list we need to move all current values to the side
func (list *List) Unshift(value interface{}) {
	newMemory := []interface{}{value}

	// using the operator  ... lets you spread out multiple values
	list.memory = append(newMemory, list.memory...)

	list.length++
}

// Shift - remove the first value from list and move remaining values back one step
func (list *List) Shift() interface{} {
	// if the memory is empty then terminate function
	if list.length == 0 {
		return nil
	}

	// decrease length and shrink memory by one, also return shifted value
	firstIndexValue := list.memory[0]
	list.memory = list.memory[1:]

	list.length--

	return firstIndexValue
}

// Get - Get will give us a specific value based on its index
func (list *List) Get(index int) interface{} {
	return list.memory[index]
}
