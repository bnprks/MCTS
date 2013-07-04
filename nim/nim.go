package nim

import (
	"fmt"
	"github.com/bnprks/MCTS/mcts"
	"math/rand"
	"strconv"
)

const (
	k       = 10
	maxMove = 3
)

type NimState struct {
	remaining, currentPlayer int
}

func MakeNimState() NimState {
	return NimState{k, 0}
}

func (ns *NimState) Clone(mcts.State) {
	return &NimState{ns.remaining, ns.currentPlayer}
}

func (ns NimState) Clone() NimState {
	return ns
}

func (ns *NimState) AvailableMoves() []string {
	currentMax := maxMove
	if currentMax > ns.remaining {
		currentMax = ns.remaining
	}
	moves := make([]string, currentMax)
	for i, _ := range moves {
		moves[i] = fmt.Sprintf("%d", i+1)
	}
	return moves
}

func (ns NimState) AvailableMoves() []string {
	currentMax := maxMove
	if currentMax > ns.remaining {
		currentMax = ns.remaining
	}
	moves := make([]string, currentMax)
	for i, _ := range moves {
		moves[i] = fmt.Sprintf("%d", i+1)
	}
	return moves
}

func (ns *NimState) MakeMove(s string) {
	ns.currentPlayer = 1 - ns.currentPlayer
	amnt, _ := strconv.ParseInt(s, 10, 0)
	ns.remaining -= int(amnt)
}

func (ns NimState) MakeMove(s string) NimState {
	ns.currentPlayer = 1 - ns.currentPlayer
	amnt, _ := strconv.ParseInt(s, 10, 0)
	ns.remaining -= int(amnt)
	return ns
}

func (ns *NimState) RandomPlayout() bool {
	player := ns.currentPlayer

	var moves []string
	var move string
	for ns.remaining > 0 {
		moves = ns.AvailableMoves()
		move = moves[rand.Intn(len(moves))]
		fmt.Println(move, ns.currentPlayer)
		ns.MakeMove(move)
	}
	//Note to self, to win player != current Player
	return player != ns.currentPlayer
}

func (ns NimState) RandomPlayout() bool {
	player := ns.currentPlayer
	var moves []string
	var move string
	for ns.remaining > 0 {
		moves = ns.AvailableMoves()
		move = moves[rand.Intn(len(moves))]
		fmt.Println(move, ns.currentPlayer)
		ns.MakeMove(move)
	}
	//Note to self, to win player != current Player
	return player != ns.currentPlayer
}
