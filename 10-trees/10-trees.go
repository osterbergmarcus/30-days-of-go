/*
* Trees are nonlinear and are used for representing data in a hierarchical structure
* Tree consists of a root node, parent nodes and child nodes
 */

package Trees

// Tree ...
type Tree struct {
	rootNode *Node
}

// Node ...
type Node struct {
	value    interface{}
	children []Node
}

// Traverse ...
func (tree *Tree) Traverse(rootNode *Node, newNode *Node, nodeValue interface{}) {

	if rootNode.value == nodeValue {
		rootNode.children = append(rootNode.children, *newNode)
		return
	}

	// for each child node traverse its children
	for i := range rootNode.children {
		tree.Traverse(&rootNode.children[i], newNode, nodeValue)
	}
}

// Add ...
func (tree *Tree) Add(data interface{}, nodeValue interface{}) {
	newNode := &Node{
		data,
		make([]Node, 0),
	}

	if tree.rootNode.value == nil {
		tree.rootNode = newNode
		return
	}

	if tree.rootNode.value == nodeValue {
		tree.rootNode.children = append(tree.rootNode.children, *newNode)
		return
	}

	// start walking the tree to find the node
	tree.Traverse(tree.rootNode, newNode, nodeValue)
}
