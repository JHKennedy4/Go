package main

type Cell struct {
	X int
	Y int
}

type Board struct {
	Cells []Cell
}

func main() {
	board := Board{[]Cell{
		Cell{2, 4},
		Cell{2, 4},
		Cell{2, 4},
	}}

	for _, cell := range board.Cells {

		neighbors := cell.GetNeighbors()
		count := board.GetLivingNeighbors(neighbors)

		if count > 3 || count < 2 {
			board.Kill(cell)
		}

	}
}

func (board *Board) IndexOf(cell Cell) int {
	for i, boardCell := range board.Cells {
		if boardCell.X == cell.X && boardCell.Y == cell.Y {
			return i
		}
	}
	return -1
}

func (board *Board) Kill(cell Cell) {
	cells := board.Cells
	i := board.IndexOf(cell)
	cells[i] = cells[len(cells)-1]
	board.Cells = cells[:len(cells)-1]
}

func (board *Board) GetLivingNeighbors(neighbors [8]Cell) int {
	for _, neighborCell := range neighbors {

	}

	return 1
}

func IsAlive(cell Cell) bool {
	return
}

func (cell *Cell) GetNeighbors() [8]Cell {
	var neighbors [8]Cell
	neighborIndex := 0
	for i := cell.X - 1; i <= cell.X+1; i++ {
		for j := cell.Y - 1; j <= cell.Y+1; j++ {
			neighborCell := Cell{i, j}
			if *cell != neighborCell {
				neighbors[neighborIndex] = Cell{i, j}
			}
		}
	}
	return neighbors
}
