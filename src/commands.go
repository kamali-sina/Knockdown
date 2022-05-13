package main

import (
	"fmt"
	"os"
)

const (
	SpacesInBetween = 5
)

type CommandHandler func([]string, *Player)

func HandleBlock(args []string, player *Player) {
	fmt.Println("handling block")
}

func HandleDodge(args []string, player *Player) {
	fmt.Println("handling dodge")
}

func HandleFight(args []string, player *Player) {
	fmt.Println("handling fight")
}

func HandleLevelup(args []string, player *Player) {
	fmt.Println("handling fight")
}

func HandleExit(args []string, player *Player) {
	fmt.Println("See you back in the ring...")
	os.Exit(0)
}

func SetActiveCommandSet(activeCommandset Commandset) {
	ActiveCommandset = activeCommandset
}
