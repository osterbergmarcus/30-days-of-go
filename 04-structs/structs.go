// https://golang.org/ref/spec#types
// https://golang.org/ref/spec#Composite_literals

package main

import (
	"fmt"
	"strings"
)

/*
* structs are used when we need a collection of data types
* we can define names and types as a fields that we then can point to
 */
type Languages struct {
	x string
	y string
}

// custom data types
type MyCustomType map[string]string

type LanguagesRecord struct {
	languages MyCustomType
}

// attatch a method to the Language struct, remember that functions takes a "receiver"?
func (l Languages) MyMethod() string {
	str := strings.ToUpper(l.x) + strings.ToUpper(l.y)
	return str
}

/*
* go has no classes and does not support inheritance
* instead go uses composition
 */
type Animal struct {
	name string
}

type Robot struct {
	Animal
}

func (f Animal) speak() {
	fmt.Println("Go!")
}

func (r Robot) talk() {
	fmt.Println("Beep")
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

	langToUpperCase := Languages{
		"go",
		"javascript",
	}

	// calling our method
	fmt.Println(langToUpperCase.MyMethod())

	// compose Robot and Animal structs
	cyborg := Robot{Animal{"gopher"}}
	cyborg.speak()
	cyborg.talk()
}
