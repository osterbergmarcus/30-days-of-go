// https://golang.org/ref/spec#types
// https://golang.org/ref/spec#Composite_literals

package main

import "fmt"

// structs are used when we need a collection of data types
// we can define the names and types as a field

type Languages struct {
	x string
	y string
}

// we can also create our custom data types
type MyCustomType map[string]string

type LanguagesRecord struct {
	languages MyCustomType
}

func main() {
	// assign values of a type using composite literals
	concurrentLanguages := Languages{
		"go",
		"elixir",
	}

	fmt.Printf("concurrent languages: %v %v\n", concurrentLanguages.x, concurrentLanguages.y)

	myLanguages := LanguagesRecord{
		map[string]string{
			"go":      "currently learning",
			"clojure": "next to learn",
		},
	}

	fmt.Println(myLanguages.languages)
}
