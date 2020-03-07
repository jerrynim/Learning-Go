package myDictionary

import "errors"

// Dictionary type
type Dictionary map[string]string

var errNotFount = errors.New("error")

//Search for a word
func (d Dictionary) Search(word string) (string, error) {
	value, exists := d[word]
	if exists {
		return value, nil
	}
	return "", errNotFount
}
