package main

import (
	"fmt"
)

func main() {
	p := MakePlayer(ReadLine())
	fmt.Println(p.getInfo())
	SetActiveCommandSet(DefendingMoves, DefendingMoveNames)
	run(p)
}
