package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strings"
	"unicode"
)

const (
	SpacesInBetween = 5
)

// TODO: ADD getch type input

func ReadLine(message string) string {
	if message != "" {
		fmt.Println(message)
	}
	reader := bufio.NewReader(os.Stdin)
	fmt.Print("> ")
	name, _ := reader.ReadString('\n')
	name = strings.Replace(name, "\n", "", -1)
	return name
}

func ReadCharNumber() int {
	reader := bufio.NewReader(os.Stdin)
	fmt.Print("> ")
	for {
		char, _, err := reader.ReadRune()

		if err != nil {
			fmt.Println(err)
		}
		if unicode.IsNumber(char) {
			return int(char - '0')
		}
	}
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
