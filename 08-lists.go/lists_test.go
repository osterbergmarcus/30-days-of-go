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

func TestGet(t *testing.T) {
	assert.Equal(t, list.Get(1), "Nice")
}
