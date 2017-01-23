// https://godoc.org/testing

package unitTests

import "testing"

func TestMyFunction(t *testing.T) {
	str := MyFunction()
	t.Log(str)
}
