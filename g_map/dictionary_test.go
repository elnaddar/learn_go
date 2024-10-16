package maps

import "testing"

func TestDictionary(t *testing.T) {
	t.Run("test search", func(t *testing.T) {
		dictionary := Dictionary{"test": "this is just a test"}
		got, _ := dictionary.Search("test")
		want := "this is just a test"

		assertStrings(t, got, want)
	})

	t.Run("unknown word", func(t *testing.T) {
		dictionary := Dictionary{"test": "this is just a test"}
		_, err := dictionary.Search("unknown")
		if err == nil {
			t.Fatal("expected to get an error.")
		}

		assertErrors(t, err, ErrNotFound)
	})
}

func TestAdd(t *testing.T) {
	t.Run("new word", func(t *testing.T) {
		word := "test"
		definition := "this is just a test"
		dictionary := Dictionary{}
		err := dictionary.Add(word, definition)
		want := "this is just a test"

		assertErrors(t, err, nil)
		assertDefinition(t, dictionary, word, want)
	})

	t.Run("existing word", func(t *testing.T) {
		word := "test"
		definition := "this is just a test"
		dictionary := Dictionary{word: definition}
		err := dictionary.Add(word, "a new test")

		assertErrors(t, err, ErrWordExists)
		assertDefinition(t, dictionary, word, definition)
	})
}

func TestUpdate(t *testing.T) {
	t.Run("update exiting word", func(t *testing.T) {
		word := "test"
		definition := "this is just a test"
		dictionary := Dictionary{word: definition}
		newDefinition := "new definition"

		err := dictionary.Update(word, newDefinition)
		assertErrors(t, err, nil)
		assertDefinition(t, dictionary, word, newDefinition)
	})

	t.Run("update word doesn't exist", func(t *testing.T) {
		dictionary := Dictionary{}
		word := "test"
		definition := "new definition"

		err := dictionary.Update(word, definition)
		assertErrors(t, err, ErrWordDoesNotExist)
	})
}

func TestDelete(t *testing.T) {
	word := "test"
	dictionary := Dictionary{word: "test definition"}

	dictionary.Delete(word)

	_, err := dictionary.Search(word)
	assertErrors(t, err, ErrNotFound)
}

func assertDefinition(t *testing.T, dictionary Dictionary, word string, definition string) {
	got, err := dictionary.Search(word)

	if err != nil {
		t.Fatal("should find added word:", err)
	}
	assertStrings(t, got, definition)
}

func assertStrings(t testing.TB, got string, want string) {
	t.Helper()
	if got != want {
		t.Errorf("got %q want %q given %q", got, want, "test")
	}
}

func assertErrors(t testing.TB, got, want error) {
	t.Helper()
	if got != want {
		t.Errorf("got %q want %q given %q", got, want, "test")
	}
}
