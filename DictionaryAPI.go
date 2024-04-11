package DictionaryAPI

import (
	"errors"
)

var (
	wordNotFound = errors.New("key not found")
	wordFound    = errors.New("word is in the dictionary, try update")
)

type Dictionary map[string]string

func (dictionary Dictionary) Search(word string) (string, error) {
	switch dictionary[word] {
	case "":
		return "", wordNotFound
	default:
		return dictionary[word], nil
	}
}

func (dictionary Dictionary) Add(word, description string) error {

	_, err := dictionary.Search(word)

	switch err {
	case wordNotFound:
		dictionary[word] = description
	case nil:
		return wordFound
	default:
		return err
	}
	return nil
}

func (dictionary Dictionary) Delete(word string) error {
	_, err := dictionary.Search(word)

	switch err {
	case wordNotFound:
		return err
	case nil:
		delete(dictionary, word)
	default:
		return err
	}
	return nil
}

func (dictionary Dictionary) Update(word, newDescription string) error {
	_, err := dictionary.Search(word)

	switch err {
	case wordNotFound:
		return err
	case nil:
		dictionary[word] = newDescription
	default:
		return err
	}
	return nil
}
