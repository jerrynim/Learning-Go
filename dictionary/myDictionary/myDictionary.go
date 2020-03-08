package myDictionary

import "errors"

// Dictionary type
type Dictionary map[string]string

var errNotFount = errors.New("error")
var errWordExists = errors.New("That word already exists")

//Search for a word
func (d Dictionary) Search(word string) (string, error) {
	value, exists := d[word]
	if exists {
		return value, nil
	}
	return "", errNotFount
}

func (d Dictionary) AddWord(word string, definition string) error {
	if _, err := d.Search(word); err == nil {
		return errWordExists
	}
	d[word] = definition
	return nil
}
