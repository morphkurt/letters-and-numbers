package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"

	"github.com/morphkurt/letters-and-numbers/letters"
	"github.com/morphkurt/letters-and-numbers/numbers"
)

var usage = `NAME:
	Countdown - Solve the countdown puzzle
	 
USAGE:
	countdown -l abcdefgh words/words.txt
	countdown -n 1,2,3,4,5,6 356`

func main() {
	s, r := run(os.Args)
	if r == 1 {
		fmt.Print(s)
		os.Exit(r)
	} else {
		fmt.Print(s)
		os.Exit(r)
	}
}

func run(args []string) (string, int) {
	argsWithoutProg := args[1:]
	if len(argsWithoutProg) == 0 {
		return usage, 1
	}
	key := argsWithoutProg[0]
	if key == "-l" {
		word := argsWithoutProg[1]
		f := argsWithoutProg[2]
		list, e := letters.FindWord(word, f)
		if e != nil {
			return "dictionary file doesn't exist", 1
		}
		if len(list) > 0 {
			r := []string{}
			r = append(r, fmt.Sprintf("Found %d words with %d letters of length\r\n", len(list), len(list[0])))
			for i, v := range list {
				r = append(r, fmt.Sprintf("%d: %s\r\n", i+1, v))
			}
			return strings.Join(r, "\n"), 0

		} else {
			return "No words found", 0
		}
	} else if key == "-n" {
		bricksValues := argsWithoutProg[1]
		numValue := strings.Split(bricksValues, ",")
		targetVal := argsWithoutProg[2]
		target, e := strconv.Atoi(targetVal)
		if e != nil {
			return usage, 1
		}
		nums := []int{}
		for _, v := range numValue {
			n, e := strconv.Atoi(v)
			if e != nil {
				return usage, 1
			}
			nums = append(nums, n)
		}
		r, p := numbers.Solve(nums, target)
		if r == target {
			return fmt.Sprintf("Target %d reached with the solve %s\n", r, p), 0
		} else {
			return fmt.Sprintf("Target %d not reached, but closest is %d with the solve %s\n", target, r, p), 0
		}
	}
	return usage, 1

}
