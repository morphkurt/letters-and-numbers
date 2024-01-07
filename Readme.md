# Word Finder in Go

This repository contains a command-line Go application that helps in finding words.

## Overview

This application parses a text file that contains a list of words and finds words that can be formed from the given input string. The application is case-sensitive and doesn't allow characters that aren't present in the input string. 

The result returns all the possible words that can be formed in the descending order of their lengths. If a word is present in the text file and is creatable from the input string, it will be included in the result. In case no words are found, generic message conveying the same will be printed.

## Files and Functions

The main file is `main.go`, which consists of `main` function, the `FindWord` function that finds all possible words from the input string.

The `readFile` function reads a text file that contains a list of words and converts its contents to a string. The path to the text file is hardcoded as `words/words.txt`.

## Usage

You can run the application using the following command:

    go run main.go <your_word>

Replace `<your_word>` with your input string.

For instance:

    go run main.go hello

This will display all the words that can be formed from the string 'hello' and are present in the words.txt file.