package main

import (
	"bytes"
	"strings"
	"testing"
	"unicode"
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

var solvedPuzzles = []string{
	`1 2 3 4 5 6 7 8 9
7 8 9 1 2 3 4 5 6 
4 5 6 7 8 9 1 2 3
9 1 2 3 4 5 6 7 8
6 7 8 9 1 2 3 4 5
3 4 5 6 7 8 9 1 2
8 9 1 2 3 4 5 6 7
5 6 7 8 9 1 2 3 4
2 3 4 5 6 7 8 9 1
`,
}

func stripSpaces(s string) string {
	return strings.Map(func(r rune) rune {
		if unicode.IsSpace(r) {
			return -1
		}
		return r
	}, s)
}

func TestSolvedPuzzle(t *testing.T) {
	for i, s := range testPuzzles {
		r := strings.NewReader(s)
		p, err := NewPuzzle(r)
		if err != nil {
			t.Errorf("error reading puzzle: %s", err.Error())
		}
		if bytes.Compare([]byte(p.String()), []byte(s)) != 0 {
			t.Errorf("puzzle in `\n%v` != puzzle out: `\n%v`", []byte(p.String()), []byte(s))
		}
		solved := p.Solve()
		if stripSpaces(solved.String()) != stripSpaces(solvedPuzzles[i]) {
			t.Errorf("solution: `\n%s` != solved: `\n%s`", solvedPuzzles[i], solved.String())
		}
	}
}
