// https://golang.org/ref/spec#Interface_types
// https://golang.org/pkg/fmt/

/*
* Functions with a signature of an interface will accept any value of any type that implements the interface
* let's implement our own interface
 */
package interfaces

import (
	"fmt"
)

type Character string

type Person string

type Response struct {
	data string
}

// function getData can change its bahaviour depending on its data type
func (c Character) requestQuote() string {
	return fmt.Sprintf("%v -%v", "Never send a human to do a machine's job", c)
}

func (p Person) requestQuote() string {
	return fmt.Sprintf("%v -%v", "The one exclusive sign of thorough knowledge is the power of teaching", p)
}

/*
* interfaces can define what actions our types can execute
* an empty interface can hold values of any type
*
* interfaces is a collection of method signatures
* it has a name and type field similiar to structs
* basiclly we're creating an interface for a method that can be
* shared amongst data types that implements the same method
 */
type Quote interface {
	requestQuote() string
}

// canTakeManyTypes can take many types and change its behaviour
func canTakeManyTypes(i Quote) map[string]string {
	if i != nil {
		return map[string]string{
			"quote": i.requestQuote(),
		}
	}
	return nil
}
