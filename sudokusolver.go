package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"unicode"
)

type Puzzle []byte

const PuzzleSize = byte(9)

// Square is a single entry in the puzzle
type Square struct {
	pos      byte
	possible []byte
}

func (s Square) Row() byte {
	return s.pos / PuzzleSize
}

func (s Square) Column() byte {
	return s.pos % PuzzleSize
}

// Box is the index of the 3x3 box this square is in; numbered top-to-bottom, left-to-right
func (s Square) Box() byte {
	return s.Column()/3 + s.Row()/3*3
}

func NewPuzzle(r io.Reader) (Puzzle, error) {
	scanner := bufio.NewScanner(r)

	var p Puzzle
	for scanner.Scan() {
		for _, c := range scanner.Bytes() {
			switch {
			case c == '_':
				p = append(p, 0)
			case unicode.IsDigit(rune(c)):
				p = append(p, byte(c-'0'))
			default:
				continue
			}
		}
	}
	if byte(len(p)) != PuzzleSize*PuzzleSize {
		return p, fmt.Errorf("expected %d squares; got %d", PuzzleSize*PuzzleSize, len(p))
	}
	return p, nil
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
