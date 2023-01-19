package dict

import "errors"

// Dictionary type (alias of type map[string]string)
// type can have methods in Go!
// by default, hashmap in Go is a reference type
// so you don't need to pass a pointer like (d *Dictionary)
type Dictionary map[string]string

var (
	errNotFound   = errors.New("Not found")
	errCantUpdate = errors.New("Can't update non-existing word")
	errCantDelete = errors.New("Can't delete non-existing word")
	errWordExists = errors.New("That word already exists")
)

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

// Update a word
func (d Dictionary) Update(word, definition string) error {
	_, err := d.Search(word)
	if err == errNotFound {
		return errCantUpdate
	}
	d[word] = definition
	return nil
}

// Delete a word
func (d Dictionary) Delete(word string) error {
	// delete doesn't return anything and does nothing if the word doesn't exist
	// which is why we need to check it if we want to know if it was deleted
	_, err := d.Search(word)
	if err == nil {
		delete(d, word)
		return nil
	}
	return errCantDelete
}
