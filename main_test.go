package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestFindWord1(t *testing.T) {
	var w = "ndbiaefrh"

	got := FindWord(w)
	want := []string{"brained", "inbread", "handier", "badiner"}

	assert.Equal(t, want, got)
}

func TestFindWord2(t *testing.T) {
	var w = "ctalekbir"

	got := FindWord(w)
	want := []string{"crablike", "atrickle"}

	assert.Equal(t, want, got)
}
