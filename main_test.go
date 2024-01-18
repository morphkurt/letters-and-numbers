package main

import (
	"testing"

	"github.com/morphkurt/letters-and-numbers/letters"
	"github.com/morphkurt/letters-and-numbers/numbers"
	"github.com/stretchr/testify/assert"
)

func TestFindWord(t *testing.T) {
	word := "abcdefg"
	f := "words/words.txt"
	list := letters.FindWord(word, f)
	assert.NotEmpty(t, list, "Expecting a non-empty list of words")
	assert.NotEqual(t, 0, len(list), "Expecting more than 0 words found")
}

func TestSolve(t *testing.T) {
	numValue := []int{1, 2, 3, 4, 5, 6}
	target := 356
	r, p := numbers.Solve(numValue, target)
	assert.NotNil(t, r, "Expecting a result number")
	assert.NotNil(t, p, "Expecting a solve pattern")
}
