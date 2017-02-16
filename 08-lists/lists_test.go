package list

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

// create a memory pointer to a new List instance
var list = &List{}

func TestPush(t *testing.T) {
	list.Push("Very")
	list.Push("Nice")
	list.Push("List")

	if len(list.memory) != 3 {
		t.Errorf("Expected length of memory to be 3 received %v", len(list.memory))
	}

	assert.Equal(t, list.memory[1], "Nice")
}

func TestPop(t *testing.T) {
	assert.Equal(t, list.Pop(), "List")

	if len(list.memory) != 2 {
		t.Errorf("Expected length of memory to be 2 received %v", len(list.memory))
	}
}

func TestUnshift(t *testing.T) {
	list.Unshift("List")

	assert.Equal(t, list.memory[0], "List")

	if len(list.memory) != 3 {
		t.Errorf("Expected length of memory to be 3 received %v", len(list.memory))
	}
}

func TestShift(t *testing.T) {
	list.Shift()

	assert.Equal(t, list.memory[0], "Very")

	if len(list.memory) != 2 {
		t.Errorf("Expected length of memory to be 2 received %v", len(list.memory))
	}
}

func TestGet(t *testing.T) {
	assert.Equal(t, list.Get(1), "Nice")
}
