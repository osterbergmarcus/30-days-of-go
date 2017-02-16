// https://godoc.org/testing
// https://github.com/stretchr/testify

package unitTests

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestMyFunction(t *testing.T) {
	str := MyFunction()

	// error will be printed only if the test fails
	t.Error(str)
}

func TestMyOtherFunction(t *testing.T) {
	assert.Equal(t, MyOtherFunction(10), 100, "Should return 100")
}
