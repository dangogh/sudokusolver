package main

import (
	"strings"
	"testing"
)

var testPuzzles = []string{
	`1 _ 3 _ 5 _ 7 _ 9
7 _ 9 _ 2 _ 4 _ 6
_ 5 _ 7 _ 9 _ 2 _
9 _ 2 _ 4 _ 6 _ 8
_ 7 _ 9 _ 2 _ 4 _
3 _ 5 _ 7 _ 9 _ 2
_ 9 _ 2 _ 4 _ 6 _
5 _ 7 _ 9 _ 2 _ 4
_ 3 _ 5 _ 7 _ 9 _
`,
}

func TestReadPuzzle(t *testing.T) {
	for _, s := range testPuzzles {
		r := strings.NewReader(s)
		p, err := NewPuzzle(r)
		if err != nil {
			t.Errorf("error reading puzzle: %s", err.Error())
		}
		if p.String() != s {
			t.Errorf("puzzle in `%s` != puzzle out: `%s`", s, p.String())
		}
	}
}
