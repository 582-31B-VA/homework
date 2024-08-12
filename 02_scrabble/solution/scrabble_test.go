package main

import "testing"

func TestLetterValue(t *testing.T) {
	tests := []struct {
		got  string
		want int
	}{
		{got: "a", want: 1},
		{got: "A", want: 1},
	}
	for _, test := range tests {
		if letterValue(test.got) != test.want {
			t.Errorf("got: %v, want: %v", test.got, test.want)
		}
	}
}

func TestWordValue(t *testing.T) {
	got := wordValue("cabbage")
	want := 14
	if got != want {
		t.Errorf("got: %v, want: %v", got, want)
	}
}
