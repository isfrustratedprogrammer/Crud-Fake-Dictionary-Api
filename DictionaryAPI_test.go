package DictionaryAPI

import (
	"testing"
)

func TestAdd(t *testing.T) {
	dictionary := Dictionary{}

	t.Run("Add word in dictionary", func(t *testing.T) {
		err := dictionary.Add("Black", "Color")

		if err != nil {
			t.Fatal(err)
		}
	})
	t.Run("Word is in dictionary", func(t *testing.T) {
		dictionary.Add("Black", "Color")

		err := dictionary.Add("Black", "Color2")

		if err == nil {
			t.Fatal(err)
		}
	})
}

func TestSearch(t *testing.T) {
	dictionary := Dictionary{"test": "testing description"}

	t.Run("Word is in dictionary", func(t *testing.T) {
		got, _ := dictionary.Search("test")
		want := "testing description"

		if got != want {
			t.Errorf("got %q want %q", got, want)
		}
	})
	t.Run("Word not in dictionary", func(t *testing.T) {
		_, err := dictionary.Search("me")

		if err == nil {
			t.Fatal(err)
		}
	})
}

func TestDelete(t *testing.T) {
	dictionary := Dictionary{"test": "testing description"}

	t.Run("Word is in dictionary", func(t *testing.T) {

		err := dictionary.Delete("test")

		if err != nil {
			t.Fatal(err)
		}
	})
	t.Run("Word not in dictionary", func(t *testing.T) {
		dictionary.Delete("123123")
	})
}

func TestUpdate(t *testing.T) {
	dictionary := Dictionary{"test": "testing description"}

	t.Run("Word is in dictionary", func(t *testing.T) {
		err := dictionary.Update("test", "testosterone")

		if err != nil {
			t.Fatal(err)
		}
	})
	t.Run("Word is not in dictionary", func(t *testing.T) {
		err := dictionary.Update("testedcase123", "testosterone")

		if err == nil {
			t.Fatal(err)
		}
	})
}
