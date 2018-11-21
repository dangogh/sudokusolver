package main

import "testing"

func TestSquare(t *testing.T) {

	var row, col, box byte
	for pos := byte(0); pos < PuzzleSize*PuzzleSize; pos++ {
		s := Square{pos: pos}
		row = pos / PuzzleSize
		col = pos % PuzzleSize
		box = 3*row + col
		if row != s.Row() {
			t.Errorf("pos %d row expected %d, got %d", pos, row, s.Row())
			t.Errorf("pos %d col expected %d, got %d", pos, col, s.Column())
			t.Errorf("pos %d box expected %d, got %d", pos, box, s.Box())
		}
	}
}
