package interfaces

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestInterface(t *testing.T) {
	var characterQuote Character = "Aristotle"
	var personQuote Person = "Agent Smith"

	assert.Equal(
		t,
		canTakeManyTypes(characterQuote),
		map[string]string{
			"quote": fmt.Sprintf("%v -%v", "Never send a human to do a machine's job", characterQuote),
		},
	)

	assert.Equal(
		t,
		canTakeManyTypes(personQuote),
		map[string]string{
			"quote": fmt.Sprintf("%v -%v", "The one exclusive sign of thorough knowledge is the power of teaching", personQuote),
		},
	)
}
