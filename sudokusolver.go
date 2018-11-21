package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strconv"
	"strings"
	"unicode"
)

type Cell struct {
	Value byte
	Pos   byte
}

func NewCell(pos, val byte) *Cell {
	return &Cell{Pos: pos, Value: val}
}

func (c Cell) Row() byte {
	return c.Pos / 9
}

func (c Cell) Column() byte {
	return c.Pos % 9
}

func (c Cell) Box() byte {
	return 3*(c.Row()/3) + c.Column()/3
}

func (c Cell) Available() []byte {
	r := c.Row().Available()
	c := c.Column().Available()
	b := c.Box().Available()
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

func (g Group) Taken() map[byte]struct{} {
	taken := make(map[byte]struct{}, 9)
	for _, c := range g {
		if c.Value == 0 {
			continue
		}
		taken[c.Value] = struct{}{}
	}
	return taken
}

type byAvail []Group

func (g byAvail) Len() int {
	return len(g)
}

func (g byAvail) Less(a, b Group) bool {
	return len(a.Taken()) > len(b.Taken())
}

func (g *byAvail) Swap(i, j int) {
	s := []Group(*g)
	s[i], s[j] = s[j], s[i]
}

func (c Cell) Taken() map[byte]struct{} {

	taken := c.Row().Taken()
	for t := range c.Column().Taken() {
		taken[t] = struct{}{}
	}
	for t := range c.Box().Taken() {
		taken[t] = struct{}{}
	}
	return taken
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
	}
	if byte(len(p)) != PuzzleSize*PuzzleSize {
		return p, fmt.Errorf("expected %d squares; got %d", PuzzleSize*PuzzleSize, len(p))
	}
	return p, nil
}

func (p Puzzle) Cells() []*Cell {
	var cells []*Cell
	for _, r := range p.Rows {
		cells = append(cells, []*Cell(r))
}

func (p Puzzle) String() string {
	var s []byte
	for i, c := range p {
		if byte(c) == 0 {
			s = append(s, '_')
		} else {
			s = append(s, byte(c)+'0')
		}
		if byte(i+1)%PuzzleSize == 0 {
			s = append(s, '\n')
		} else {
			s = append(s, ' ')
		}
	}
	return cells
}

func (p Puzzle) Groups() []Group {
	groups := append(p.Rows, p.Columns...)
	return append(groups, p.Boxes...)
}

func (p Puzzle) Solve() Puzzle {
	fmt.Printf("puzzle is\n%v\n", p)
	cells := p.Cells()
	for i, c := range cells {
		if c == nil {
			continue
		}
		taken := c.Taken()
		switch len(taken) {
		case 8:
			// one number left..
			var j byte
			for j = 1; j <= 9; j++ {
				if _, ok := taken[j]; !ok {
					c.Fill(j)
					taken[j] = struct{}{}
					break
				}
			}
		}
		if len(taken) == 9
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
