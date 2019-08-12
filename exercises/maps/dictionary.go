package dictionary

import "errors"

// ErrNotFound ...
var ErrNotFound = errors.New("could not find the word you were looking for")

// Dictionary ...
type Dictionary map[string]string

// Add ..
func (d Dictionary) Add(word, definition string) {
	d[word] = definition
}

// Search ..
// Map is reference by type, so no need to pass pointer
// Maps being a reference is really good,
// because no matter how big a map gets there will only be one copy.
func (d Dictionary) Search(word string) (string, error) {
	definition, ok := d[word]
	if !ok {
		return "", ErrNotFound
	}

	return definition, nil
}

func dictionary() {

}
