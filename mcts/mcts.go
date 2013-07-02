package mcts

import (
	"math"
)

const (
	ExplorationFactor = 1.0 //Constant to multiply the UCB term
	MinVisits         = 1   //The minimum number of visits before a node gets expanded
)

type node struct {
	//Win and visit counts are float64s for convencience during calculations, but should remain integers
	wins       float64
	visits     float64
	children   []node //I hope these don't need to be pointers
	parent     *node
	isLeaf     bool
	isTerminal bool
	move       string //The move that was made to get to this position from the parent
}

type State interface {
	Clone() State             // copies all the values from the argument to the callee, to make a copy without allocation
	AvailableMoves() []string //List the legal moves from the current state in string format
	MakeMove(string)          //Make a move given a string (of the same format that AvailableMoves outputs)

	//Make a random playout from the current state and return whether the first player to move won 
	//(counting from the beginning of the random playout)
	RandomPlayout() bool
}

func (n *node) selectChild() *node {
	var max float64 = 0.0 //since all UCT values are positive, we can use 0 as the max value
	var bestChild node

	var UCB float64
	for _, child := range n.children {
		UCB = child.averageScore() + ExplorationFactor*math.Sqrt(2.0*math.Log(n.visits)/child.visits)
		if UCB > max {
			bestChild = child
			max = UCB
		}
	}
	return &bestChild
}

func (n *node) averageScore() float64 {
	return n.wins / n.visits
}

func (n *node) expand(s State) {
	n.isLeaf = false
	moves := s.AvailableMoves()
	l := len(moves)
	if len(moves) == 0 {
		//Works on the assumption that a lack of remaining moves signifies an ended game
		n.isTerminal = true
	}
	n.children = make([]node, l)
	for i, child := range n.children {
		child.move = moves[i]
		child.isLeaf = true
		child.parent = n
	}
}

func (n *node) treePolicy(s State) *node {
	//S is a state consistant with node n
	ret := n
	for !ret.isLeaf && !ret.isTerminal {
		ret.visits++
		ret = ret.selectChild()
		s.MakeMove(ret.move)
	}
	if ret.visits > MinVisits {
		ret.expand(s)
		ret = ret.selectChild()
		s.MakeMove(ret.move)
	}
	return ret
}

func (n *node) backPropogate(result bool) {
	for currNode := n; currNode != nil; currNode = currNode.parent {
		if result {
			currNode.wins++
		}
		result = !result
	}
}

func MonteCarlo(s State, times int) string {
	var root node
	root.expand(s)
	var workingState State
	var leafNode *node
	var isWin bool
	for i := 0; i < times; i++ {
		workingState = s.Clone()
		leafNode = root.treePolicy(workingState)
		isWin = workingState.RandomPlayout()
		leafNode.backPropogate(isWin)
	}
	bestResult := float64(0)
	bestMove := ""
	for _, child := range root.children {
		if child.averageScore() > bestResult {
			bestResult = child.averageScore()
			bestMove = child.move
		}
	}
	return bestMove
}
