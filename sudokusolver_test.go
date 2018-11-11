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

var solvedPuzzles = []string{
	`
1 2 3 4 5 6 7 8 9
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

func TestSolvedPuzzle(t *testing.T) {
	for i, s := range testPuzzles {
		r := strings.NewReader(s)
		p, err := NewPuzzle(r)
		if err != nil {
			t.Errorf("error reading puzzle: %s", err.Error())
		}
		if p.String() != s {
			t.Errorf("puzzle in `\n%s` != puzzle out: `\n%s`", s, p.String())
		}

		// each group has 9 cells, each cell has 3 groups
		groupCount := make(map[*Cell]int, 9*9)
		for _, g := range p.Groups() {
			if len(g) != 9 {
				t.Errorf("expected group %d to have 9 cells; got %d", i, len(g))
			}
			for _, c := range g {
				if _, ok := groupCount[c]; !ok {
					groupCount[c] = 1
				} else {
					groupCount[c]++
				}
			}
		}
		if len(groupCount) != 9*9 {
			t.Errorf("expected %d cells, found %d", 9*9, len(groupCount))
		}
		for _, n := range groupCount {
			if n != 3 {
				t.Errorf("cell should have 3 groups; found %d", n)
			}
		}

		solved := p.Solve()
		if solved.String() != solvedPuzzles[i] {
			t.Errorf("solution: `\n%s` != solved: `\n%s`", solvedPuzzles[i], solved.String())
		}
	}
}
