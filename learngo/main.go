package main

import (
	"fmt"

	"github.com/salientsoph/learngo/mydict"
)

func main() {
	dictionary := mydict.Dictionary{"first": "First word"}
	/*dictionary["hello"] = "hello"
	fmt.Println(dictionary)*/
	fmt.Println(dictionary["first"])

	definition, err := dictionary.Search("second")

	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(definition)
	}

}
