package numbers

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSolve1(t *testing.T) {
	numbers := []int{25, 50, 75, 2, 6, 8}
	want := 856
	got, gotPath := Solve(numbers, 856)
	wantPath := "75+25*8+50+6"
	assert.Equal(t, want, got)
	assert.Equal(t, wantPath, gotPath)
}

func TestSolve2(t *testing.T) {
	numbers := []int{25, 50, 75, 2, 6, 8}
	want := 856
	got, gotPath := Solve(numbers, 855)
	wantPath := "50-2*6*75/25-8"
	assert.Equal(t, want, got)
	assert.Equal(t, wantPath, gotPath)
}
