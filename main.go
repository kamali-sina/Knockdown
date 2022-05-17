package main

import (
	"fmt"
)

func main() {
	InitEnemyRegistery()
	fmt.Println(EnemyRegistery)
	x := MakeEnemy("BackstreetFighter")
	fmt.Println(x)
	p := MakePlayer(ReadLine())
	fmt.Println(p.getInfo())
	SetActiveCommandSet(MainMenuCommandset)
	run(p)
}
