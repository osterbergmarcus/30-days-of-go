package functions

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestMultipleReturns(t *testing.T) {
	lang, learning := MultipleReturns("Go")
	t.Logf("%v %t\n", lang, learning)

	fmt.Printf("%v %t\n", lang, learning)
}

func TestFactorial(t *testing.T) {
	// 5!
	val := Factorial(5)
	t.Log(val)

	assert.Equal(t, val, 120, "result should equal 120")

	fmt.Println("result should equal 120", val == 120)
}

func TestCurry(t *testing.T) {
	val := Curry("Learning ")("Golang")
	t.Log(val)

	fmt.Println(val)
}
