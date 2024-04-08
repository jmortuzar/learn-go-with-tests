package main

import "testing"

const thisIsJustATest = "this is just a test"

func TestSearch(t *testing.T) {
	dictionary := Dictionary{"test": thisIsJustATest}

	t.Run("known word", func(t *testing.T) {
		got, _ := dictionary.Search("test")
		want := thisIsJustATest
		asserStrings(t, got, want)
	})

	t.Run("unknown word", func(t *testing.T) {
		_, err := dictionary.Search("unknown")

		assertError(t, err, ErrNotFound)
	})
}

func asserStrings(t testing.TB, got, want string) {
	t.Helper()
	if got != want {
		t.Errorf("got %q want %q given %q", got, want, "test")
	}
}

func assertError(t testing.TB, got, want error) {
	t.Helper()
	if got != want {
		t.Errorf("got %q want %q", got, want)
	}
}

func TestAdd(t *testing.T) {
	t.Run("new word", func(t *testing.T) {
		definition := thisIsJustATest
		word := "test"
		dictionary := Dictionary{word: definition}
		err := dictionary.Add(word, "new test")

		assertError(t, err, ErrWordExists)
		assertDefinition(t, dictionary, word, definition)
	})

	t.Run("existing word", func(t *testing.T) {
		definition := thisIsJustATest
		word := "test"
		dictionary := Dictionary{}
		dictionary.Add(word, definition)

		assertDefinition(t, dictionary, word, definition)
	})

}

func assertDefinition(t testing.TB, dictionary Dictionary, word, definition string) {
	t.Helper()
	got, err := dictionary.Search(word)

	if err != nil {
		t.Fatal("should find added word:", err)
	}

	asserStrings(t, got, definition)
}

func TestUpdate(t *testing.T) {
	t.Run("existing word", func(t *testing.T) {
		definition := thisIsJustATest
		word := "test"
		dictionary := Dictionary{word: definition}
		newDefinition := "new definition"

		err := dictionary.Update(word, newDefinition)

		assertError(t, err, nil)
		assertDefinition(t, dictionary, word, newDefinition)
	})

	t.Run("new word", func(t *testing.T) {
		definition := thisIsJustATest
		word := "test"
		dictionary := Dictionary{}

		err := dictionary.Update(word, definition)

		assertError(t, err, ErrWordDoesNotExist)
	})

}

func TestDelete(t *testing.T) {
	definition := thisIsJustATest
	word := "test"
	dictionary := Dictionary{word: definition}

	dictionary.Delete(word)

	_, err := dictionary.Search(word)
	if err != ErrNotFound {
		t.Errorf("expected %q to be deleted", word)
	}
}
