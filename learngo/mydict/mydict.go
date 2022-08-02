package mydict

import "errors"

// Dictionary == map[string]string
// Dictionary type
type Dictionary map[string]string

var errNotFound = errors.New("Not found")

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
