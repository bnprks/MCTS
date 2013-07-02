package main

import (
	"fmt"
	"github.com/bnprks/MCTS/mcts"
	"github.com/bnprks/MCTS/nim"
)

func main() {
	fmt.Println("hello world")
	s := nim.MakeNimState()
	str := mcts.MonteCarlo(&s, 1)
	fmt.Println(str)
}
