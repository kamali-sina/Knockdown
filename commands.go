package main

import (
	"fmt"
	"sort"
	"strings"
)

const (
	SpacesInBetween = 5
)

type CommandHandler func([]string, *Player)

var ActiveCommandset map[int]CommandHandler
var ActiveCommandsetNames map[int]string

var DefendingMoveNames = map[int]string{
	1: "Block",
	2: "Dodge",
}

var DefendingMoves = map[int]CommandHandler{
	1: HandleBlock,
	2: HandleDodge,
}

func HandleBlock(args []string, player *Player) {
	fmt.Println("handling block")
}

func HandleDodge(args []string, player *Player) {
	fmt.Println("handling dodge")
}

func getCommandRowLenght() int {
	maxLen := -1
	for _, val := range ActiveCommandsetNames {
		x := len(val)
		if x > maxLen {
			maxLen = x
		}
	}
	return 3 + SpacesInBetween + maxLen + 3
}

func PrintAvailableCommands() {
	rowLen := getCommandRowLenght()
	fmt.Println(strings.Repeat("-", rowLen))
	keys := make([]int, 0)
	for k := range ActiveCommandsetNames {
		keys = append(keys, k)
	}
	sort.Ints(keys)
	for _, key := range keys {
		fmt.Printf("|%2d.%v%v |\n", key, strings.Repeat(" ", SpacesInBetween), ActiveCommandsetNames[key])
	}
	fmt.Println(strings.Repeat("-", rowLen))
}

func SetActiveCommandSet(activeCommandset map[int]CommandHandler, activeCommandsetNames map[int]string) {
	ActiveCommandset = activeCommandset
	ActiveCommandsetNames = activeCommandsetNames
}
