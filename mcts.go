package mcts

import (
	"fmt"
	"math/rand"
	"math"
	"log"
)

const (
	ExplorationFactor = 1.0 //Constant to multiply the UCB term
	MinVisits = 1 //The minimum number of visits before a node gets expanded
)

type node struct {
	//Win and visit counts are float64s for convencience during calculations, but should remain integers
	wins float64
	visits float64
	children []node //I hope these don't need to be pointers
	parent *node
	isLeaf bool
	move string //The move that was made to get to this position from the parent
}

type state interface {
	Clone(state) // copies all the values from the argument to the callee, to make a copy without allocation
	ListAvaliableMoves() []string //List the legal moves from the current state in string format
	MakeMove(string) //Make a move given a string (of the same format that ListAvaliableMoves outputs)

	//Make a random playout from the current state and return whether the first player to move won 
	//(counting from the beginning of the random playout)
	RandomPlayout() bool 
}

func (n *node) selectChild() *node {
	max := 0 //since all UCT values are positive, we can use 0 as the max value
	var bestChild node

	var UCTScore float64
	for _, child := range n.children {
		UCT = child.averageScore() + ExplorationFactor * math.Sqrt(2.0 * math.Log(n.visits)/ child.visits)
		if(UCT > max) {
			bestChild = child
			max = UCT
		}
	}
	return &bestChild
}

func (n *node) averageScore() float64 {
	return n.wins/n.visits
}

func (n *node) expand(s state) {
	n.isLeaf = false
	moves := s.ListAvaliableMoves()
	l := len(moves)
	n.childdren := make([]node, l)
	for i, child := range n.children {
		child.move = moves[i]
		child.isLeaf = true
		child.parent = &n
	}
}

func (n *node) treePolicy(s state) *node{
	//S is a state consistant with node n
	ret := n
	for !ret.isLeaf {
		ret.visits ++
		ret = ret.selectChild()
		s.MakeMove(ret.move)
	}
	if ret.visits > MinVisits {
		ret.expand()
		ret = ret.selectChild()
		s.MakeMove(ret.move)
	}
	return ret, s
}

func (n *node) backPropogate(bool result) {
	for currNode := n; currNode != nil; currNode = currNode.parent {
		if(result) {
			currNode.wins++
		}
		result = !result
	}
}

func MonteCarlo(s state, times int) string {
	var root node
	root.expand(s)
	var workingState state
	var isWin bool
	for i := 0; i < times; i++ {
		workingState.Clone(s)
		leafState = root.treePolicy(workingState)
		isWin = workingState.RandomPlayout()
		leafState.backPropogate(isWin)
	}
	bestResult := float64(0)
	bestMove := ""
	for _, child := range root.children {
			if(child.averageScore() > bestResult) {
				bestResult = child.averageScore()
				bestMove = child.move
			}
	}
	return bestMove
}