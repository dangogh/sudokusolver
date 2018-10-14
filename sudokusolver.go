package main

import (
	"bufio"
	"errors"
	"fmt"
	"io"
	"os"
	"strconv"
	"unicode"
)

type Cell byte
type Row []Cell
type Puzzle []Row

func NewPuzzle(r io.Reader) (Puzzle, error) {
	scanner := bufio.NewScanner(r)

	var rows []Row
	for scanner.Scan() {
		var row []Cell
		for _, c := range scanner.Bytes() {
			switch {
			case c == '_':
				row = append(row, Cell(0))
			case unicode.IsDigit(rune(c)):
				row = append(row, Cell(c-'0'))
			default:
				continue
			}
		}
		if len(row) == 0 {
			continue
		}
		if len(row) != 9 {
			return Puzzle{}, errors.New("bad row with " + strconv.Itoa(len(row)) + "entries")
		}
		rows = append(rows, row)
	}
	return Puzzle(rows), nil
}

func (p Puzzle) String() string {
	var s []byte
	for _, r := range []Row(p) {
		for _, b := range r {
			if byte(b) == 0 {
				s = append(s, '_')
			} else {
				s = append(s, byte(b)+'0')
			}
			s = append(s, ' ')
		}
		s[len(s)-1] = '\n'
	}
	return string(s)
}

func (p Puzzle) Solve() Puzzle {
	fmt.Printf("puzzle is %v\n", p)
	return p
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
