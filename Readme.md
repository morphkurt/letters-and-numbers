# Countdown Puzzle Solver

This is a Go program that solves the popular game show puzzle, [Countdown](https://en.wikipedia.org/wiki/Countdown_(game_show)) and French version [Des chiffres et des lettres](https://en.wikipedia.org/wiki/Des_chiffres_et_des_lettres). It can solve both the letter and number games. 

## Prerequisites

Before running this application, ensure you have Go installed on your local machine.

## Dependencies

This project depends on two external packages:

- `github.com/morphkurt/letters-and-numbers/letters` - This package is used to process the letters segment of the countdown game.
  
- `github.com/morphkurt/letters-and-numbers/numbers` - This package is used to process the numbers segment of the countdown game.

## Usage

Clone the repository and navigate to the project directory:

```bash
git clone <repository_url>
cd <project_directory>
```

You can run the program via the terminal using the following commands:

For the letters game:

```bash
go run main.go -l abcdefgh words/words.txt
```

For the numbers game:

```bash
go run main.go -n 1,2,3,4,5,6 356
```

## Commands

The `-l` command is used to solve the letters game of the countdown. It takes a string of random letters as its first argument, and a file path to a dictionary of words as its second argument. It tries to form a word from the given letters with reference to the supplied dictionary file.

The `-n` command is used to solve the numbers game of Countdown. It takes two arguments, the first is a list of comma-separated numbers, and the second is a target number. The program tries to combine the given numbers through various mathematical operations to reach the target number.

Please note that when either of the commands are not supplied correctly, the program will provide a usage guide and terminate.