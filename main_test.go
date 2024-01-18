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
	list, _ := letters.FindWord(word, f)
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

func TestRunLetters(t *testing.T) {
	args := []string{"dummy", "-l", "abcdefg", "words/words.txt"}
	s, _ := run(args)
	expected := "Found 7 words with 5 letters of length\r\n\n1: faced\r\n\n2: begad\r\n\n3: badge\r\n\n4: debag\r\n\n5: cadge\r\n\n6: caged\r\n\n7: fadge\r\n"
	assert.Equal(t, expected, s)
}

func TestRunLettersWithOneMatch(t *testing.T) {
	args := []string{"dummy", "-l", "countdown", "words/words.txt"}
	s, _ := run(args)
	expected := "Found 1 words with 9 letters of length\r\n\n1: countdown\r\n"
	assert.Equal(t, expected, s)
}

func TestRunNumbers(t *testing.T) {
	args := []string{"dummy", "-n", "1,2,3,4,5,6", "356"}
	s, _ := run(args)
	expected := "Target 356 reached with the solve 6*5*3-1*4\n"
	assert.Equal(t, expected, s)
}

func TestRunWithInvalidFlag(t *testing.T) {
	args := []string{"dummy", "-e", "1,2,3,4,5,6", "356"}
	s, _ := run(args)
	expected := "NAME:\n\tCountdown - Solve the countdown puzzle\n\t \nUSAGE:\n\tcountdown -l abcdefgh words/words.txt\n\tcountdown -n 1,2,3,4,5,6 356"
	assert.Equal(t, expected, s)
}

func TestRunWithNotEnoughArguments(t *testing.T) {
	args := []string{"dummy"}
	s, _ := run(args)
	expected := "NAME:\n\tCountdown - Solve the countdown puzzle\n\t \nUSAGE:\n\tcountdown -l abcdefgh words/words.txt\n\tcountdown -n 1,2,3,4,5,6 356"
	assert.Equal(t, expected, s)
}

func TestRunWithNumbersNotReachingTarget(t *testing.T) {
	args := []string{"dummy", "-n", "25,50,75,2,6,8", "855"}
	s, _ := run(args)
	expected := "Target 855 not reached, but closest is 856 with the solve 50-2*6*75/25-8\n"
	assert.Equal(t, expected, s)
}
