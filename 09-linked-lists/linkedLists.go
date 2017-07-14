/*
* Linked list is used as foundation for many data structures
* it's a sequence of elements known as nodes
* every node has a value and a pointer to its sibling
* Linked lists are good for adding or removing data without reallocating memory
 */

package LinkedLists

// LinkedList - our structure.
type LinkedList struct {
	head *Node // the first node in memory is refered to as "head"
	size int
}

// Node - our node will consist of a value and a pointer to the next node
// imagine the data structure like this
// [value, pointer -> [value, pointer -> [value, nil]
type Node struct {
	value interface{}
	next  *Node
}

// Size - returns the number of nodes
func (linkedList *LinkedList) Size() int {
	return linkedList.size
}

// Get - reads values from memory
func (linkedList *LinkedList) Get(position int) interface{} {
	/*
	* in a linked list data are not stored in a indexed sequence
	* so we can't randomly access the memory
	 */
	if linkedList.Size() != 0 && position <= linkedList.Size() {
		// keep track of the nodes while searching for right one
		currentNode := linkedList.head

		// jump through each node until we find the position we are looking for, starting from the head
		for i := 0; i != position; i++ {
			currentNode = currentNode.next
		}

		// finally return the current node value
		return currentNode.value
	}

	// return nil if position is out of memory range
	return nil
}

// Add - appends a node
func (linkedList *LinkedList) Add(value interface{}) {
	newNode := &Node{value: value}

	// if the linked list holds zero values then start the chain by adding a new node as head
	if linkedList.size == 0 {
		linkedList.head = newNode
		linkedList.size++
		return
	}

	// cache current node
	currentNode := linkedList.head

	// loop through the chain of nodes until current node is pointing to nil
	for currentNode.next != nil {
		currentNode = currentNode.next
	}

	// pointing current node to our new node
	currentNode.next = newNode

	// finally increase the size of the linked list
	linkedList.size++
}

// Remove - removes a node
func (linkedList *LinkedList) Remove(position int) {
	// if memory size is zero then terminate
	if linkedList.size == 0 {
		return
	}

	// if memory position don't exist then terminate
	if linkedList.Get(position) == nil {
		return
	}

	// replace head with next node
	if position == 0 {
		linkedList.head = linkedList.head.next
		return
	}

	previousNode := &Node{value: nil}
	currentNode := linkedList.head

	for i := 0; i != position; i++ {
		previousNode = currentNode
		currentNode = currentNode.next
	}

	// take the next node and point previous node to it
	previousNode.next = currentNode.next

	linkedList.size--
}

// Contains - Returns a slice of all nodes values
func (linkedList *LinkedList) Contains() []interface{} {
	values := make([]interface{}, linkedList.Size(), linkedList.Size())

	currentNode := linkedList.head
	for i := range values {
		values[i] = currentNode.value
		currentNode = currentNode.next
	}

	return values
}
