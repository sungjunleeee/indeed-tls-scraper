package dict

import "errors"

// Dictionary type (alias of type map[string]string)
// type can have methods in Go!
type Dictionary map[string]string

var errNotFound = errors.New("Not found")
var errWordExists = errors.New("That word already exists")

// Search returns the definition of a word
func (d Dictionary) Search(word string) (string, error) {
	definition, exists := d[word]
	if !exists {
		return "", errNotFound
	}
	return definition, nil
}

// Add adds a word and its definition to the dictionary
func (d Dictionary) Add(word string, def string) error {
	_, err := d.Search(word)
	if err == errNotFound {
		d[word] = def
		return nil
	}
	return errWordExists
}
