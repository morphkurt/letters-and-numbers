package letters

import (
	"errors"
	"os"
	"sort"
	"strings"
)

func readFile(f string) (string, error) {
	b, err := os.ReadFile(f) // just pass the file name
	if err != nil {
		return "", err
	}
	str := string(b) // convert content to a 'string'
	return str, nil
}

func FindWord(word, f string) ([]string, error) {
	out := []string{}
	result := []string{}
	s, err := readFile(f)
	if err != nil {
		return out, errors.New("dictonoary file doesn't exists")
	}
	lines := strings.Split(s, "\r\n")
	for _, w := range lines {
		m := copyMap(word)
		notFound := false
		for _, c := range w {
			if v, ok := m[c]; ok {
				if v == 0 {
					notFound = true
					break
				} else {
					m[c] = m[c] - 1
				}
			} else {
				notFound = true
				break
			}
		}
		if !notFound {
			out = append(out, w)
		}
	}
	sort.Slice(out, func(i, j int) bool {
		return len(out[i]) > len(out[j])
	})
	for _, w := range out {
		if len(w) >= len(out[0]) {
			result = append(result, w)
		}
	}
	return result, nil
}

func copyMap(word string) map[rune]int {
	m := map[rune]int{}
	for _, c := range word {
		if v, ok := m[c]; ok {
			m[c] = v + 1
		} else {
			m[c] = 1
		}
	}
	return m
}
