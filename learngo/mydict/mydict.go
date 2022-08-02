package mydict

import "errors"

// Dictionary == map[string]string
// Dictionary type
type Dictionary map[string]string

var (
	errNotFound   = errors.New("Not found")
	errWordExists = errors.New("That word already exists")
	errCantUpdate = errors.New("Can't update non-existing word")
)

// Search for a word
func (d Dictionary) Search(word string) (string, error) {

	/*
		i, ok := m["route"]
		-> dictionary에 해당하는 key가 존재하면 값은 value 혹은 ok
		-> ok: boolean. 값이 존재하는지 여부를 체크함
	*/
	value, exists := d[word]
	if exists {
		return value, nil
	}
	return "", errNotFound
}

// Add a word to the dictionary
func (d Dictionary) Add(word string, def string) error {
	_, err := d.Search(word) //존재하는 단어인지 체크

	/*
		if err == errNotFound {
			d[word] = def
		} else if err == nil { //존재하는 단어라면
			return errWordExists
		}
		return nil
	*/

	switch err {
	case errNotFound:
		d[word] = def
	case nil:
		return errWordExists
	}
	return nil

}

// Update a dictionary
func (d Dictionary) Update(word, definition string) error {
	_, err := d.Search(word)
	switch err {
	case nil:
		d[word] = definition
	case errNotFound:
		return errCantUpdate
	}
	return nil
}

// Delete from dictionary
func (d Dictionary) Delete(word string) {
	//이미 내재되어 있는 함수
	delete(d, word)
}
