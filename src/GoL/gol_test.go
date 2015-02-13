package main

import "testing"

/*
func TestReverse(t *testing.T) {
	cases := []struct {
		in, want string
	}{
		{"Hello, world", "dlrow ,olleH"},
		{"Hello, 世界", "界世 ,olleH"},
		{"", ""},
	}

	for _, c := range cases {
		got, _ := Reverse(c.in)
		if got != c.want {
			t.Errorf("Reverse(%q) == %q, want %q", c.in, got, c.want)
		}
	}
}
*/

func TestGetNeighbors(t *testing.T) {
	cell := Cell{2, 4}

	neighbors := cell.GetNeighbors()

	if len(neighbors) != 8 {
		t.Errorf("Expected 8 neighbors: found %d", neighbors)
	}

}

func TestKill(t *testing.T) {
	board := Board{[]Cell{
		Cell{2, 4},
		Cell{3, 4},
		Cell{4, 4},
	}}

	boardLength := len(board.Cells)

	board.Kill(Cell{2, 4})

	if len(board.Cells) != boardLength-1 {

		t.Errorf("Expected the newBoard to have smae length as boad after kill: found %+v", board)
	}
}

func TestGetLivingNeighbors(t *testing.T) {
	board := Board{[]Cell{
		Cell{2, 4},
		Cell{3, 4},
		Cell{4, 4},
	}}

	cell := Cell{2, 4}

	neighbors := cell.GetNeighbors()

	count := board.GetLivingNeighbors(neighbors)

}
