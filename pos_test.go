package main

import "testing"

func TestCell(t *testing.T) {

	for pos := 0; pos < PuzzleSize*PuzzleSize; pos++ {
		s := Cell{Pos: pos}
		row := pos / PuzzleSize
		col := pos % PuzzleSize
		box := 3*row + col
		if row != s.Row() {
			t.Errorf("pos %d row expected %d, got %d", pos, row, s.Row())
		}
		if col != s.Column() {
			t.Errorf("pos %d col expected %d, got %d", pos, col, s.Column())
		}
		if box != s.Box() {
			t.Errorf("pos %d box expected %d, got %d", pos, box, s.Box())
		}
	}
}
