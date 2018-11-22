package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strings"
	"unicode"
)

const PuzzleSize = 9

type Cell struct {
	Value byte
	Pos   int
}

func (c Cell) Row() int {
	return c.Pos / 9
}

func (c Cell) Column() int {
	return c.Pos % 9
}

func (c Cell) Box() int {
	return 3*(c.Row()/3) + c.Column()/3
}

func (c *Cell) Fill(b byte) {
	c.Value = b
}

func (c Cell) String() string {
	i := c.Value
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

func (p Puzzle) Taken(c Cell) map[byte]struct{} {
	taken := p.Rows[c.Row()].Taken()
	for t := range p.Columns[c.Column()].Taken() {
		taken[t] = struct{}{}
	}
	for t := range p.Boxes[c.Box()].Taken() {
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

	var pos int
	for scanner.Scan() {
		for _, b := range scanner.Bytes() {
			var val byte
			switch {
			case b == '_':
				val = 0
			case unicode.IsDigit(rune(b)):
				val = b - '0'
			default:
				continue
			}
			c := &Cell{Pos: pos, Value: val}
			fmt.Printf("got a cell{pos:%d,val:%d}: %++v\n", pos, val, c)
			p.Rows[c.Row()] = append(p.Rows[c.Row()], c)
			p.Columns[c.Column()] = append(p.Columns[c.Column()], c)
			p.Boxes[c.Box()] = append(p.Boxes[c.Box()], c)
			pos++
		}
	}
	if pos != PuzzleSize*PuzzleSize {
		return p, fmt.Errorf("expected %d squares; got %d", PuzzleSize*PuzzleSize, pos)
	}
	return p, nil
}

func (p Puzzle) Cells() []*Cell {
	var cells []*Cell
	for _, r := range p.Rows {
		cells = append(cells, []*Cell(r)...)
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
		taken := p.Taken(*c)
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
			cells[i] = nil
		case 9:
			cells[i] = nil
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
