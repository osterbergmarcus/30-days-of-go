// https://godoc.org/testing
// https://github.com/stretchr/testify

package unitTests

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestMyFunction(t *testing.T) {
	str := MyFunction()
	t.Log(str)
}

func TestMyOtherFunction(t *testing.T) {
	assert.Equal(t, MyOtherFunction(10), 100, "Should return 100")
}
