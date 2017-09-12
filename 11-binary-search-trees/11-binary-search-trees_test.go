package BST

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

// create a memory pointer to a new instance
var tree = new(BST)

/*
 *create a tree structure as illustrated
 *         5
 *        / \
 *       4   6
 *      /\   /\
 *     3       7
 *             /\
 *              100
 *               /\
 *              99 101
 */

func TestAdd(t *testing.T) {
	tree.Insert(5)
	tree.Insert(4)
	tree.Insert(6)
	tree.Insert(3)
	tree.Insert(7)
	tree.Insert(100)
	tree.Insert(99)
	tree.Insert(101)

	// assert on left side of the tree
	assert.Equal(t, tree.rootNode.value, 5)
	assert.Equal(t, tree.rootNode.left.value, 4)
	assert.Equal(t, tree.rootNode.left.left.value, 3)

	// assert on right side of the tree
	assert.Equal(t, tree.rootNode.right.value, 6)
	assert.Equal(t, tree.rootNode.right.right.value, 7)
	assert.Equal(t, tree.rootNode.right.right.right.value, 100)
	assert.Equal(t, tree.rootNode.right.right.right.left.value, 99)
	assert.Equal(t, tree.rootNode.right.right.right.right.value, 101)
}

func TestFind(t *testing.T) {
	// contains node
	assert.Equal(t, tree.Find(101), true)

	// not containing node
	assert.Equal(t, tree.Find(1), false)
}
