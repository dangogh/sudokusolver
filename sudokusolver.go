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

func (c Cell) String() string {
	i := byte(c)
	if i == 0 {
		return "_"
	}
	return string(i + byte('0'))
}

type Group []*Cell

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

func (p Puzzle) Solve() Puzzle {
	fmt.Printf("puzzle is\n%v\n", p)
	groups := append(p.Rows, p.Columns...)
	groups = append(groups, p.Boxes...)

	/*
		minAvail := 9
		changed := false
			// sort by number of available numbers
			for _, g := range groups {
				a := g.Available()
				switch len(a) {
				case 0:
					continue
				case 1:
					for _, c := range g {
						if c.Value() == 0 {
							c.Fill(a[0])
							break
						}
					}
				default:
					if len(a) < len(minGroup.Available()) {
						minGroup = g
					}
				}
			}
	*/
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
