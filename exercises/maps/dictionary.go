package dictionary

// var ErrNotFound = errors.New("could not find the word you were looking for")

const (
	// ErrNotFound ...
	ErrNotFound = DictionaryErr("could not find the word you were looking for")
	// ErrWordExists ..
	ErrWordExists = DictionaryErr("cannot add word because it already exists")
	// ErrWordDoesNotExist ..
	ErrWordDoesNotExist = DictionaryErr("cannot update word because it does not exist")
)

// DictionaryErr ...
type DictionaryErr string

func (e DictionaryErr) Error() string {
	return string(e)
}

// Dictionary ...
type Dictionary map[string]string

// Add ..
func (d Dictionary) Add(word, definition string) error {
	_, err := d.Search(word)

	switch err {
	case ErrNotFound:
		d[word] = definition
	case nil:
		return ErrWordExists
	default:
		return err
	}

	return nil
}

// Update ...
func (d Dictionary) Update(word, definition string) error {
	_, err := d.Search(word)

	switch err {
	case ErrNotFound:
		return ErrWordDoesNotExist
	case nil:
		d[word] = definition
	default:
		return err
	}

	return nil
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

// Delete ...
func (d Dictionary) Delete(word string) error {
	_, err := d.Search(word)

	switch err {
	case ErrNotFound:
		return ErrWordDoesNotExist
	case nil:
		delete(d, word)
	default:
		return err
	}

	return nil
}

func dictionary() {

}

/**

A gotcha that reference types introduce is that maps can be a nil value.
A nil map behaves like an empty map when reading,
but attempts to write to a nil map will cause a runtime panic

Therefore, you should never initialize an empty map variable:


DON'T

var m map[string]string

DO

var dictionary = map[string]string{}

// OR

var dictionary = make(map[string]string)

*/
