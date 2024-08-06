# Homework 2: Scrabble

Scrabble is a word game where players place letter tiles on a board to
form words. Each letter has the following value.

| Letter              | Value  |
| ------------------- | ------ |
| a e i o u l n r s t | 1      |
| d g                 | 2      |
| b c m p             | 3      |
| f h v w y           | 4      |
| k                   | 5      |
| j x                 | 8      |
| q z                 | 10     |

A word's score is the sum of its letters' values. For example, the word
'cabbage' is worth 14 points: 3 points for 'c', 1 point for 'a', 3
points for 'b', 3 points for 'b', 1 point for 'a', 2 points for 'g', and
1 point for 'e'.

Write a command-line application that prints a given word's value. Users
may enter words with capital letters. You are not allowed to use
`fmt.Scan` or similar functions. Instead, your application must accept
the word as an argument.

```sh
go run scrabble.go cabbage
```

The `os.Args` variable from the [os][] package is an array of strings
that holds the command-line arguments, starting with the program name.

[os]: https://pkg.go.dev/os
