package LinkedLists

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

// create a memory pointer to a new instance
var linkedList = &LinkedList{}

type gopher struct {
	gopher string
}

func TestAdd(t *testing.T) {
	linkedList.Add(gopher{"oNe"})
	linkedList.Add(gopher{"TwO"})
	linkedList.Add(gopher{"tHrEe"})
	linkedList.Add(gopher{"fOUR"})
	linkedList.Add(gopher{"Five"})

	if linkedList.Size() != 5 {
		t.Errorf("Expected size of memory to be 5 received %v", linkedList.Size())
	}
}

func TestGet(t *testing.T) {
	node := linkedList.Get(3)

	assert.Equal(t, node, gopher{"fOUR"})
}

func TestRemove(t *testing.T) {
	linkedList.Remove(3)

	if linkedList.Size() != 4 {
		t.Errorf("Expected size of memory to be 4 received %v", linkedList.Size())
	}
}

func TestContains(t *testing.T) {
	values := linkedList.Contains()
	gophers := make([]interface{}, linkedList.Size(), linkedList.Size())

	gophers[0] = gopher{"oNe"}
	gophers[1] = gopher{"TwO"}
	gophers[2] = gopher{"tHrEe"}
	gophers[3] = gopher{"Five"}

	assert.Equal(t, values, gophers)
}
