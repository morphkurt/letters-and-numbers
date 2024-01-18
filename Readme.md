[![Test Coverage](https://raw.githubusercontent.com/wiki/morphkurt/letters-and-numbers/coverage.svg)](https://raw.githack.com/wiki/morphkurt/letters-and-numbers/coverage.html)
![CI](https://github.com/morphkurt/letters-and-numbers/actions/workflows/go.yml/badge.svg)
[![Go Report Card](https://goreportcard.com/badge/github.com/morphkurt/letters-and-numbers)](https://goreportcard.com/report/github.com/morphkurt/letters-and-numbers)

# Countdown Puzzle Solver

![Logo_Des_chiffres_et_des_lettres_2016](https://github.com/morphkurt/letters-and-numbers/assets/20348847/daf32822-c6df-4fe5-bff0-67883e7f1de2)


This is a Go program that solves the popular game show puzzle, [Countdown](https://en.wikipedia.org/wiki/Countdown_(game_show)) / French version [Des chiffres et des lettres](https://en.wikipedia.org/wiki/Des_chiffres_et_des_lettres) and Letters and Numbers Australian version. It can solve both the letter and number games. 

## Prerequisites


Before running this application, ensure you have Go installed on your local machine.

## Dependencies

This project depends on two external packages:

- `github.com/morphkurt/letters-and-numbers/letters` - This package is used to process the letters segment of the countdown game.
  
- `github.com/morphkurt/letters-and-numbers/numbers` - This package is used to process the numbers segment of the countdown game.

## Usage

Clone the repository and navigate to the project directory:

```bash
git clone https://github.com/morphkurt/letters-and-numbers.git
cd letters-and-numbers
```

You can run the program via the terminal using the following commands:

For the letters game:

```bash
go run main.go -l abcdefgh words/words.txt
Found 1 words with 9 letters of length
1: countdown
```

For the numbers game:

```bash
go run main.go -n 1,2,3,4,5,6 356
Target 356 reached with the solve 6*5*3-1*4
```

## Commands

The `-l` command is used to solve the letters game of the countdown. It takes a string of random letters as its first argument, and a file path to a dictionary of words as its second argument. It tries to form a word from the given letters with reference to the supplied dictionary file.

The `-n` command is used to solve the numbers game of Countdown. It takes two arguments, the first is a list of comma-separated numbers, and the second is a target number. The program tries to combine the given numbers through various mathematical operations to reach the target number.

Please note that when either of the commands are not supplied correctly, the program will provide a usage guide and terminate.
