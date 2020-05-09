package main

import (
	"fmt"

	"github.com/jerrynim/learngo/dictionary/myDictionary"
)

func main() {
	dictionary := myDictionary.Dictionary{}
	word := "hello"
	definition := "Greeting"
	err := dictionary.AddWord(word, definition)
	if err != nil {
		fmt.Println(err)
	}
	searchedValue, _ := dictionary.Search(word)
	fmt.Println("found", word, "definition:", searchedValue)
	addingError := dictionary.AddWord(word, definition)
	if addingError != nil {
		fmt.Println(addingError)
	}
	dictionary.Delete("hello")
}
