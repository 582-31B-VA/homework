package main

import (
	"fmt"
	"os"
	"slices"
	"strings"
)

// letterValue returns the value of the letter l.
func letterValue(l string) int {
	l = strings.ToLower(l)
	letterValues := map[int][]string{
		1:  {"a", "e", "i", "o", "u", "l", "n", "r", "s", "t"},
		2:  {"d", "g"},
		3:  {"b", "c", "m", "p"},
		4:  {"f", "h", "v", "w", "y"},
		5:  {"k"},
		8:  {"j", "x"},
		10: {"q", "z"},
	}
	for value, letters := range letterValues {
		if slices.Contains(letters, l) {
			return value
		}
	}
	return 0
}

// wordValue returns the value of the word w.
func wordValue(w string) int {
	if len(w) == 0 {
		return 0
	}
	letters := strings.Split(w, "")
	return letterValue(letters[0]) + wordValue(w[1:])
}

func main() {
	word := os.Args[1]
	value := wordValue(word)
	fmt.Println(value)
}
