package main

import (
	"fmt"
)

func main() {
	p := MakePlayer(ReadLine())
	fmt.Println(p.getInfo())
	SetActiveCommandSet(MainMenuCommandset)
	run(p)
}
