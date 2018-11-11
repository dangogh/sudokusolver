package main

import (
	"bufio"
	"errors"
	"fmt"
	"io"
	"os"
	"strconv"
	"strings"
	"unicode"
)

type Cell byte

func NewCell(val byte) *Cell {
	cell := Cell(val)
	return &cell
}

func (c Cell) Value() byte {
	return byte(c)
}

func (c *Cell) Fill(b byte) {
	*c = Cell(b)
}

func (c Cell) String() string {
	i := byte(c)
	if i == 0 {
		return "_"
	}
	return string(i + byte('0'))
}

type Group []*Cell

func (g Group) Available() []byte {
	taken := make(map[byte]struct{}, 9)
	for _, c := range g {
		if c.Value() == 0 {
			continue
		}
		taken[c.Value()] = struct{}{}
	}
	var avail []byte
	var v byte
	for v = 0; v < 9; v++ {
		if _, ok := taken[v]; !ok {
			avail = append(avail, v)
		}
	}
	return avail
}

type byAvail []Group

func (g byAvail) Len() int {
	return len(g)
}

func (g byAvail) Less(a, b Group) bool {
	return len(a.Available()) < len(b.Available())
}

func (g *byAvail) Swap(i, j int) {
	s := []Group(*g)
	s[i], s[j] = s[j], s[i]
}

type Puzzle struct {
	Rows    []Group
	Columns []Group
	Boxes   []Group
}

func NewPuzzle(r io.Reader) (Puzzle, error) {
	scanner := bufio.NewScanner(r)

	var p Puzzle
	p.Rows = make([]Group, 9)
	p.Columns = make([]Group, 9)
	p.Boxes = make([]Group, 9)

	var rowIdx int
	for scanner.Scan() {
		var row []*Cell
		for _, b := range scanner.Bytes() {
			var c *Cell
			switch {
			case b == '_':
				c = NewCell(0)
			case unicode.IsDigit(rune(b)):
				c = NewCell(b - '0')
			default:
				continue
			}
			row = append(row, c)
		}
		if len(row) == 0 {
			continue
		}
		if len(row) != 9 {
			return Puzzle{}, errors.New("bad row with " + strconv.Itoa(len(row)) + "entries")
		}
		p.Rows[rowIdx] = Group(row)
		for colIdx, c := range row {
			p.Columns[colIdx] = append(p.Columns[colIdx], c)
			boxIdx := 3*(rowIdx/3) + colIdx/3
			p.Boxes[boxIdx] = append(p.Boxes[boxIdx], c)
		}
		rowIdx++
	}
	return p, nil
}

func (p Puzzle) Groups() []Group {
	groups := append(p.Rows, p.Columns...)
	return append(groups, p.Boxes...)
}

func (p Puzzle) Solve() Puzzle {
	fmt.Printf("puzzle is\n%v\n", p)
	minGroup := Group{}
	changed := true
	// sort by number of available numbers
	for changed {
		changed = false
		for _, g := range p.Groups() {
			a := g.Available()
			switch len(a) {
			case 0:
				continue
			case 1:
				for _, c := range g {
					if c.Value() == 0 {
						c.Fill(a[0])
						changed = true
						break
					}
				}
			default:
				if len(a) < len(minGroup.Available()) {
					minGroup = g
				}
			}
		}
	}
	return p
}

func (p Puzzle) String() string {
	var s string
	for _, r := range p.Rows {
		var ss []string
		for _, c := range r {
			ss = append(ss, c.String())
		}
		s += strings.Join(ss, " ") + "\n"
	}
	return s
}

func main() {
	for _, fn := range os.Args[1:] {
		f, err := os.Open(fn)
		if err != nil {
			panic(err.Error())
		}
		puzzle, err := NewPuzzle(f)
		puzzle.Solve()
		fmt.Println(puzzle.String())
	}

}
