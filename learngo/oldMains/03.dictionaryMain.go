//package main
package oldMains

import (
	"fmt"

	"github.com/salientsoph/learngo/mydict"
)

func dictionaryMain() {
	dictionary := mydict.Dictionary{"first": "First word"}
	/*dictionary["hello"] = "hello"
	fmt.Println(dictionary)*/
	fmt.Println(dictionary["first"])

	/* search */
	definition, err := dictionary.Search("second")

	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(definition)
	}

	/* add */
	word := "hello"
	defi := "Greeting"

	addErr := dictionary.Add(word, defi)
	if addErr != nil {
		fmt.Println(addErr)
	}

	searchDefinition, _ := dictionary.Search(word)
	fmt.Println("found ", "'", word, "'", "definition: ", searchDefinition)

	addErr2 := dictionary.Add(word, defi)
	if addErr2 != nil {
		fmt.Println(addErr2)
	}

	/* update */
	updateWord := "hello"
	updateErr := dictionary.Update(updateWord, "second")
	if updateErr != nil {
		fmt.Println(updateErr)
	}
	updateWord2, _ := dictionary.Search(updateWord)
	fmt.Println("definition for updateWord: ", updateWord2)

	/* delete */
	dictionary.Delete(word)
	fmt.Println(dictionary)
}
