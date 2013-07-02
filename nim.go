package nim

import (
	"fmt"
	"math/rand"
	"strconv"
)		

const(
	k := 10
	maxMove := 3
)

type NimState struct {
	remaining, currentPlayer int
}	

func (ns *NimState) Clone(other *NimState) {
	ns.remaining = other.remaining
	ns.currentPlayer = other.currentPlayer
}

func (ns *NimState) AvailableMoves() []string {
	currentMax := maxMove
	if currentMax > ns.remaining {
		currentMax = ns.remaining
	}
	moves = make([]string, maxMove - 1)
	for i, _ := range moves {
		moves[i] = fmt.Sprintf("%d", i + 1)
	}
	return moves
}

func (ns *NimState) MakeMove(s string) {
	ns.currentPlayer = 1 - ns.curentPlayer
	amnt, _ := strconv.ParseInt(s, 10, 0)
	ns.remaining -= amnt
}

func (ns *NimState) RandomPlayout() bool {
	player := ns.currentPlayer
	
	moves []string
	move string
	for ns.remaining > 0 {
		moves = ns.AvailableMoves()
		move = moves[rand.Intn(len(moves))]
		ns.MakeMove(move)
	}
	//Note to self, to win player != current Player
	return player != ns.currentPlayer
}