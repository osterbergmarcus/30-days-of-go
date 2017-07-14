package structs

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestLanguages(t *testing.T) {
	// assign values of a type using composite literals
	concurrentLanguages := Languages{
		"go",
		"elixir",
	}

	assert.Equal(t, concurrentLanguages.x, "go")
	assert.Equal(t, concurrentLanguages.y, "elixir")
}

func TestLanguagesRecord(t *testing.T) {
	myLanguages := LanguagesRecord{
		map[string]string{
			"go":      "currently learning",
			"clojure": "next to learn",
		},
	}

	assert.Equal(t, myLanguages.languages["go"], "currently learning")
	assert.Equal(t, myLanguages.languages["clojure"], "next to learn")
}

func TestMyMethod(t *testing.T) {
	langToUpperCase := Languages{
		"go",
		"javascript",
	}

	assert.Equal(t, langToUpperCase.MyMethod(), "GOJAVASCRIPT", "Method should return strings uppercased")
}

func TestComposingStructs(t *testing.T) {
	// compose Robot and Animal structs
	cyborg := Robot{Animal{"gopher"}}

	assert.Equal(t, cyborg.speak(), "Go!")
	assert.Equal(t, cyborg.talk(), "Beep")
}
