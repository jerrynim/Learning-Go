package myDictionary

import "errors"

// Dictionary type
type Dictionary map[string]string

var errNotFount = errors.New("error")
var errWordExists = errors.New("That word already exists")
var errCantUpdate = errors.New("cant update does not exist")

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

func (d Dictionary) Update(word, definition string) error {
	_, err := d.Search(word)
	switch err {
	case nil:
		d[word] = definition
	case errNotFount:
		return errCantUpdate
	}
	return nil
}

func (d Dictionary) Delete(word string) {
	_, err := d.Search(word)
	if err == nil {
		delete(d, word)
	}
}
