/*
* Binary search trees divides your data in half on every traverse. This makes
* searching for values really performant. Data flows in a "unidirectional" way
* this means that we can't connect any child nodes directly to each other
 */

package BST

// BST ...
type BST struct {
	rootNode *Node
}

// Node ...
type Node struct {
	value int
	left  *Node
	right *Node
}

// Adds a node in a sorted order
func (tree *BST) Insert(value int) {
	newNode := &Node{value: value}

	// just in case root is not defined
	if tree.rootNode == nil {
		tree.rootNode = newNode
		return
	}

	// keep track of current position in the tree
	currentNode := tree.rootNode

	for {
		if value > currentNode.value {
			// if right leaf node is not defined then add the new node and stop
			// the traversing
			if currentNode.right == nil {
				currentNode.right = newNode
				break
			}
			// otherwise continue traversing the right leaf nodes
			currentNode = currentNode.right
		} else if value < currentNode.value {
			// again if left leaf node is not defined then add the new node and stop
			// the traversing
			if currentNode.left == nil {
				currentNode.left = newNode
				break
			}
			// otherwise continue traversing the left leaf nodes
			currentNode = currentNode.left
		} else {
			// if we land here that means that value has already been added to
			// the tree. Stop the travese
			return
		}
	}
}

// BST structures are known for fast lookup time
func (tree *BST) Find(value int) bool {
	currentNode := tree.rootNode

	// Pretty much the same operation for searching nodes as adding nodes in
	// the tree. However instead of assigining nodes we want to return true or
	// false if given node is found
	for {
		if value > currentNode.value {
			if currentNode.right == nil {
				return false
			}
			currentNode = currentNode.right
		} else if value < currentNode.value {
			if currentNode.left == nil {
				return false
			}
			currentNode = currentNode.left
		} else {

			// if value is neither less or greater then any of the traversed
			// nodes, then we successfully found what we were looking for
			return true
		}
	}
}
