// https://golang.org/ref/spec#types
// https://golang.org/ref/spec#Composite_literals

package main

import (
	"fmt"
	"strings"
)

// structs are used when we need a collection of data types
// we can define names and types as a fields
// that we then later can point to
type Languages struct {
	x string
	y string
}

// we can also create our custom data types
type MyCustomType map[string]string

type LanguagesRecord struct {
	languages MyCustomType
}

// lets attatch a method to our Language struct
// remember that functions takes a "receiver"?
func (l Languages) MyMethod() string {
	str := strings.ToUpper(l.x) + strings.ToUpper(l.y)
	return str
}

func main() {
	// assign values of a type using composite literals
	concurrentLanguages := Languages{
		"go",
		"elixir",
	}

	fmt.Printf("concurrent languages: %v %v\n", concurrentLanguages.x, concurrentLanguages.y)

	// use our own type
	myLanguages := LanguagesRecord{
		map[string]string{
			"go":      "currently learning",
			"clojure": "next to learn",
		},
	}

	fmt.Println(myLanguages.languages)

	// Calling our method
	langToUpperCase := Languages{
		"go",
		"javascript",
	}

	fmt.Println(langToUpperCase.MyMethod())
}
