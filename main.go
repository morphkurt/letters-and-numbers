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
	argsWithoutProg := os.Args[1:]
	if len(argsWithoutProg) == 0 {
		fmt.Println(usage)
		os.Exit(1)
	}
	key := argsWithoutProg[0]
	if key == "-l" {
		word := argsWithoutProg[1]
		f := argsWithoutProg[2]
		list := letters.FindWord(word, f)
		if len(list) > 0 {
			fmt.Printf("Found %d words with %d letters of length\r\n", len(list), len(list[0]))
			for i, v := range list {
				fmt.Printf("%d: %s\r\n", i+1, v)
			}
		} else {
			fmt.Printf("No words found")
		}
	} else if key == "-n" {
		bricksValues := argsWithoutProg[1]
		numValue := strings.Split(bricksValues, ",")
		targetVal := argsWithoutProg[2]
		target, e := strconv.Atoi(targetVal)
		if e != nil {
			fmt.Println(usage)
			os.Exit(1)
		}
		nums := []int{}
		for _, v := range numValue {
			n, e := strconv.Atoi(v)
			if e != nil {
				fmt.Println(usage)
				os.Exit(1)
			}
			nums = append(nums, n)
		}
		r, p := numbers.Solve(nums, target)
		if r == target {
			fmt.Printf("Target %d reached with the solve %s\n", r, p)
		} else {
			fmt.Printf("Target %d not reached, but closest is %d with the solve %s\n", target, r, p)
		}
	} else {
		fmt.Println(usage)
		os.Exit(1)
	}
}
