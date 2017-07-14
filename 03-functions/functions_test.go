package functions

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestMultipleReturns(t *testing.T) {
	lang, learning := MultipleReturns("Go")
	assert.Equal(t, lang, "Learning Go")
	assert.Equal(t, learning, true)
}

func TestFactorial(t *testing.T) {
	// 5!
	val := Factorial(5)

	assert.Equal(t, val, 120, "should equal 120")
}

func TestCurry(t *testing.T) {
	val := Curry("Learning ")("Golang")
	assert.Equal(t, val, "Learning Golang")

}
