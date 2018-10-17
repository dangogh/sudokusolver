package main

import "testing"

func TestSquare(t *testing.T) {

	var row = 1
	var col = 0
	var box = 0
	for pos := 0; pos < 9*9; pos++ {
		s := Square{pos: pos}
		row = pos / 9
		col = pos % 9
		box = 3*row + col
		if row != s.Row() {
			t.Errorf("pos %d row expected %d, got %d", pos, row, s.Row())
			t.Errorf("pos %d col expected %d, got %d", pos, col, s.Column())
			t.Errorf("pos %d box expected %d, got %d", pos, box, s.Box())
		}
	}
}
