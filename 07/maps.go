package maps

type Dictionary map[string]string

const ErrNotFound = DictionaryErr("couldn't find the word you were looking for")
const ErrWordExists = DictionaryErr("this word is already exist")
const ErrWordDoesNotExist = DictionaryErr("this word still does not exist")

type DictionaryErr string //mivel implementalja az error interfacet, csinalhatunk belole konstanst lsd.fent

// Any type with the Error()string method fulfils the error interface
// It looks similar to the errors.errorString() implementation that powers erorrs.New()
// Howerver unlike errors.errorString() this type is a constant expression

func (e DictionaryErr) Error() string { //emiatt a DictionaryErr tipusunk implementalja az error interfacet
	return string(e)
} //https://dave.cheney.net/2016/04/07/constant-errors

func (d Dictionary) Search(word string) (string, error) {
	definition, ok := d[word]
	if !ok {
		return "", ErrNotFound
	}

	return definition, nil
}

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

// Update() takes 2 strings as arguments, which are a key and a new value for this key
func (d Dictionary) Update(word, newDefinition string) error {
	_, err := d.Search(word) //a search() altal megtalalt szo valtozojara itt nincs szuksegem, csak a hibauzenetre

	switch err { //megnezzuk milyen hibak tortenhetnek
	case ErrNotFound: //ha nem talaltuk meg a szot
		return ErrWordDoesNotExist //adja vissza, h ez a szo meg nem letezik a szotarunkban
	case nil: //ha hiba nelkul lezajlott a kereses, hajtsa vegre a frissitest
		d[word] = newDefinition //megadom a d nevu map-nak hogy a word kulcs mezonek a newDefinition stringet adja meg erteknek
	default: //alapesetben pedig (azaz minden mas esetben)
		return err //adja vissza az err valtozoban eltarolt hibauzenetet
	}

	return nil //a funkcion pedig adjon vissza egy nil erteku hibauzit ha lezajlott a frissites
}

func (d Dictionary) Delete(word string) {
	delete(d, word)
}
