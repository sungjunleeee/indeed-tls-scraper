package dict

import "errors"

// Dictionary type (alias of type map[string]string)
// type can have methods in Go!
type Dictionary map[string]string

var errNotFound = errors.New("Not found")

// Search returns the definition of a word
func (d Dictionary) Search(word string) (string, error) {
	definition, exists := d[word]
	if !exists {
		return "", errNotFound
	}
	return definition, nil
}
