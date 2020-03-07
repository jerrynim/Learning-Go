package main

import (
	"fmt"

	"github.com/jerrynim/learngo/dictionary/myDictionary"
)

func main() {
	dictionary := myDictionary.Dictionary{"first": "First word"}
	definition, err := dictionary.Search("first")
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(definition)
}
