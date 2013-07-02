package dots

import "fmt"

type Board struct {
	score, opponentScore, rows, cols int
	board                            [][]int
}

type Index struct {
	row, col int
}

func MakeBoard(rows int, cols int) Board {
	arr := make([][]int, rows)
	for i := range arr {
		arr[i] = make([]int, cols)
	}
	return Board{0, 0, rows, cols, arr}
}

func (b *Board) onBoard(idx Index) bool {
	return (0 <= idx.row && idx.row < b.rows && 0 <= idx.col && idx.col < b.cols)
}

func (b *Board) get(idx Index) *int {
	return &b.board[idx.row][idx.col]
}

func (b *Board) getNorth(idx Index) Index {
	return Index{idx.row - 1, idx.col}
}

func (b *Board) getSouth(idx Index) Index {
	return Index{idx.row + 1, idx.col}
}

func (b *Board) getWest(idx Index) Index {
	return Index{idx.row, idx.col - 1}
}

func (b *Board) getEast(idx Index) Index {
	return Index{idx.row, idx.col + 1}
}

func (b *Board) getNeighbor(idx Index, side int) Index {
	switch side {
	case 1:
		return b.getNorth(idx)
	case 2:
		return b.getEast(idx)
	case 4:
		return b.getSouth(idx)
	case 8:
		return b.getWest(idx)
	default:
		return Index{-1, -1}
	}
}

func (b *Board) move(idx Index, side int) {
	*b.get(idx) |= side
	oppositeSide := oppSide(side)
	adjacent := b.getNeighbor(idx, side)
	if b.onBoard(adjacent) {
		*b.get(adjacent) |= oppositeSide
	}
	if *b.get(idx) == 15 {
		b.score++
	} else {
		b.switchTurns()
	}
}

func (b *Board) switchTurns() {
	temp := b.score
	b.score = b.opponentScore
	b.opponentScore = temp
}

func numEdges(edges int) int {
	count := 0
	for bit := 1; bit <= 8; bit <<= 1 {
		if bit&edges != 0 {
			count++
		}
	}
	return count
}

func oppSide(side int) int {
	if side > 2 {
		return side / 4
	} else {
		return side * 4
	}
}

func (b *Board) hasEdge(idx Index, side int) bool {
	return (*b.get(idx) & side) != 0
}

func (b *Board) nextNeighbor(idx Index, side int) (adj Index, newSide int, hasErr bool) {
	if numEdges(*b.get(idx)) != 2 && numEdges(*b.get(idx)) != 3 {
		fmt.Println("had less than 2 edges")
		return idx, side, true
	}
	for i := 1; i <= 8; i <<= 1 {
		if side != i && !b.hasEdge(idx, i) {
			fmt.Println("idx", idx, "Side", side, "i", i)
			adj = b.getNeighbor(idx, i)
			newSide = oppSide(i)
			fmt.Println(adj, hasErr, newSide)
			return
		}
	}
	//If we get here then there wasn't a neighbor on the board
	fmt.Println("got to end")
	return idx, side, true
}

func (b *Board) stringLength(start Index) (count int, isLoop bool) {
	count = 0
	isLoop = true
	addToCount := func(idx Index) {
		if numEdges(*b.get(idx)) == 2 || numEdges(*b.get(idx)) == 3 {
			fmt.Println("in Count", idx)
			count++
		}
	}
	addToCount(start)
	//Finds a cycle, row or whatnot.
	//The other end of the string/cycle
	end := start
	//The side of end or start that already was seen
	var endNeighbor, startNeighbor int
	var hasErr bool
	end, endNeighbor, hasErr = b.nextNeighbor(start, 0)

	if hasErr {
		return
	}
	startNeighbor = oppSide(endNeighbor)
	for end != start && !hasErr && b.onBoard(end) {
		fmt.Println("places", end, start)
		addToCount(end)
		end, endNeighbor, hasErr = b.nextNeighbor(end, endNeighbor)
	}
	hasErr = false
	if start == end {
		return
	}
	start, startNeighbor, hasErr = b.nextNeighbor(start, startNeighbor)
	for !hasErr && b.onBoard(start) {
		fmt.Println("places", end, start)
		addToCount(start)
		start, startNeighbor, hasErr = b.nextNeighbor(start, startNeighbor)
	}
	if !b.onBoard(end) || !b.onBoard(start) {
		isLoop = false
	}
	return
}

func (b *Board) print() {
	out := ""
	for i := 0; i < b.rows; i++ {
		for j := 0; j < b.cols; j++ {
			out += "."
			if (b.hasEdge(Index{i, j}, 1)) {
				out += "-"
			} else {
				out += " "
			}
		}
		out += ".\n"
		for j := 0; j < b.cols; j++ {
			if (b.hasEdge(Index{i, j}, 8)) {
				out += "|"
			} else {
				out += " "
			}
			out += " "
		}
		if (b.hasEdge(Index{i, b.cols - 1}, 2)) {
			out += "|"
		} else {
			out += " "
		}
		out += "\n"
	}
	for j := 0; j < b.cols; j++ {
		out += "."
		if (b.hasEdge(Index{b.rows - 1, j}, 4)) {
			out += "-"
		} else {
			out += " "
		}
	}
	out += "."
	fmt.Println(out)
}

func main2() {
	b := MakeBoard(4, 4)
	fmt.Println(b.board)
	b.move(Index{0, 0}, 8)
	fmt.Println(b.board)
	b.move(Index{0, 0}, 1)
	b.move(Index{0, 1}, 1)
	b.move(Index{0, 1}, 2)
	b.move(Index{1, 1}, 2)

	b.move(Index{1, 0}, 2)
	b.move(Index{1, 0}, 8)
	b.move(Index{2, 1}, 2)
	b.move(Index{2, 1}, 4)
	b.move(Index{2, 0}, 4)
	b.move(Index{2, 0}, 8)
	b.move(Index{0, 1}, 2)
	b.move(Index{0, 1}, 4)
	fmt.Println(b.board)
	fmt.Println(b.stringLength(Index{0, 1}))
	b.print()
	fmt.Println(b.onBoard(Index{-1, -1}))

}
