package maps

import (
	"testing"
)

func TestSearch(t *testing.T) {
	dictionary := Dictionary{"test": "this is just a test"}

	t.Run("word found", func(t *testing.T) {
		got, _ := dictionary.Search("test")
		want := "this is just a test"

		assertStrings(t, got, want)
	})

	t.Run("word found, but doesn't have a description", func(t *testing.T) {
		got, _ := dictionary.Search("what")
		want := ""

		assertStrings(t, got, want)
	})

	t.Run("unknown word", func(t *testing.T) {
		_, got := dictionary.Search("unknown")

		assertError(t, got, ErrNotFound)
	})
}

func TestAdd(t *testing.T) {
	t.Run("new word", func(t *testing.T) {
		dictionary := Dictionary{}
		word := "new"
		definition := "this is a new word"

		err := dictionary.Add(word, definition)

		assertError(t, err, nil)
		assertDefinition(t, dictionary, word, definition)
	})

	t.Run("already existing word", func(t *testing.T) {
		word := "old"
		definition := "this is an old word"
		dictionary := Dictionary{word: definition}
		err := dictionary.Add(word, "new definition")

		assertError(t, err, ErrWordExists)
		assertDefinition(t, dictionary, word, definition)

	})
}

func TestUpdate(t *testing.T) {
	t.Run("existing word", func(t *testing.T) {
		word := "pulp"
		definition := "fiction"
		dictionary := Dictionary{word: definition} //feltoltom adattal a Dictionary{} mappet
		// hogy tudjunk benne frissiteni
		newDefinition := "fucking.fiction" //az uj definicio

		err := dictionary.Update(word, newDefinition)

		assertError(t, err, nil)                             //megnezem, h az update() method visszaadott error erteke nil-e(azaz hiba nelkul futott e le a frissites)
		assertDefinition(t, dictionary, word, newDefinition) //megnezem h update-eles utan benne van e dictionary mapben a word key az uj definicioval
	})

	t.Run("new word", func(t *testing.T) {
		word := "pulp"
		definition := "fiction"
		dictionary := Dictionary{}

		err := dictionary.Update(word, definition)

		assertError(t, err, ErrWordDoesNotExist) //ellenorzom, h az update() visszaadott error valtozo erteke megegyezik-e a konstans hibauzenet ertekevel,
		//tehat a frissites tenylegesen azert nem ment vegbe, mert nincs benne ez a szo a dictionary map-ben
	})
}

func TestDelete(t *testing.T) {
	word := "cinderella"
	dictionary := Dictionary{word: "tests a shoe"}

	dictionary.Delete(word)

	_, err := dictionary.Search(word)
	if err != ErrNotFound {
		t.Errorf("Expected %q to be deleted", word)
	}

}

// assertDefinition() ellenorzom, h a dictionary nevu mapben benne van a word key es a hozzatartozo definition value
func assertDefinition(t testing.TB, dictionary Dictionary, word, definition string) {
	t.Helper()

	got, err := dictionary.Search(word)

	if err != nil {
		t.Fatal("should find added word: ", err)
	}

	assertStrings(t, got, definition)
}

// assertError() ellenorzom h ket error tipusu ertek egyforma e
func assertError(t *testing.T, got, want error) {
	t.Helper()
	if got != want {
		t.Errorf("got %q want %q", got, want)
	}
}

// asserrtStrings() ellenorzom, h ket string tipusu ertek egyforma e
func assertStrings(t testing.TB, got, want string) {
	t.Helper()
	if got != want {
		t.Errorf("got %q want %q", got, want)
	}
}
