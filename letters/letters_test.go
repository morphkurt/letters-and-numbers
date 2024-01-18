package letters

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestFindWord1(t *testing.T) {
	var w = "ndbiaefrh"

	f := "../words/words.txt"
	got := FindWord(w, f)
	want := []string{"brained", "inbread", "handier", "badiner"}

	assert.Equal(t, want, got)
}

func TestFindWord2(t *testing.T) {
	var w = "ctalekbir"

	f := "../words/words.txt"
	got := FindWord(w, f)
	want := []string{"crablike", "atrickle"}

	assert.Equal(t, want, got)
}
