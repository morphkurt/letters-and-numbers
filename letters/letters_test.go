package letters

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestFindWord1(t *testing.T) {
	var w = "ndbiaefrhn"

	f := "../words/words.txt"
	got, _ := FindWord(w, f)
	want := []string{"endbrain"}

	assert.Equal(t, want, got)
}

func TestFindWord2(t *testing.T) {
	var w = "ctalekbir"

	f := "../words/words.txt"
	got, _ := FindWord(w, f)
	want := []string{"crablike", "atrickle"}

	assert.Equal(t, want, got)
}

func TestErrorIsRaisedForNonExistingFile(t *testing.T) {
	var w = "ctalekbir"

	f := "../words/non-existing-file.txt"
	_, err := FindWord(w, f)
	assert.ErrorContains(t, err, "dictonoary file doesn't exists")
}
