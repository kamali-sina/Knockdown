package main

import (
	"fmt"
)

func main() {
	InitEnemyRegistery()
	p := MakePlayer(ReadLine("Enter you name: "))
	fmt.Println(p.getInfo())
	SetActiveCommandSet(MainMenuCommandset)
	run(p)
}
