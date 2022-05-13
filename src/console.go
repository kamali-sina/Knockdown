package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strings"
)

func ReadLine() string {
	reader := bufio.NewReader(os.Stdin)
	fmt.Print("> ")
	name, _ := reader.ReadString('\n')
	name = strings.Replace(name, "\n", "", -1)
	return name
}

func getCommandRowLenght() (int, int) {
	maxLen := -1
	for _, val := range ActiveCommandset.Names {
		x := len(val)
		if x > maxLen {
			maxLen = x
		}
	}
	return 3 + SpacesInBetween + maxLen + 3, maxLen
}

func PrintAvailableCommands() {
	rowLen, maxLen := getCommandRowLenght()
	fmt.Println(strings.Repeat("-", rowLen))
	keys := make([]int, 0)
	for k := range ActiveCommandset.Names {
		keys = append(keys, k)
	}
	sort.Ints(keys)
	for _, key := range keys {
		commandName := ActiveCommandset.Names[key]
		fmt.Printf("|%2d.%v%v%v |\n", key, strings.Repeat(" ", SpacesInBetween), commandName, strings.Repeat(" ", maxLen-len(commandName)))
	}
	fmt.Println(strings.Repeat("-", rowLen))
}
