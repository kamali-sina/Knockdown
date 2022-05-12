package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
	"strconv"
)

const (
	BaseMaxHP       = 20
	BaseMaxStamina  = 50
	SaveFilePostFix = ".kds"
)

type Player struct {
	Name    string
	Hp      float64
	Stamina float64
	Xp      int
}

func (p *Player) getInfo() string {
	info := ""
	info += p.Name + "\n"
	info += "\thp:      " + strconv.FormatFloat(p.Hp, 'f', 1, 64)
	info += "\n\tstamina: " + strconv.FormatFloat(p.Stamina, 'f', 1, 64)
	info += "\n\txp:      " + strconv.Itoa(p.Xp)
	return info
}

func (p *Player) saveToFile(savefileName string) error {
	file, _ := json.MarshalIndent(*p, "", " ")

	err := ioutil.WriteFile(savefileName, file, 0644)

	return err
}

func doesFileExist(path string) bool {
	_, err := os.Stat(path)
	if err == nil {
		return true
	} else {
		return false
	}
}

func MakePlayer(name string) *Player {
	savefileName := name + SaveFilePostFix
	var player *Player
	if doesFileExist(savefileName) {
		fmt.Println("exsts!1")
		// TODO:
		// loadPlayerFromFile(savefileName)
	} else {
		player = &Player{name, BaseMaxHP, BaseMaxStamina, 0}
		player.saveToFile(savefileName)
	}

	return player
}
