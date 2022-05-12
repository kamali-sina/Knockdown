package main

import (
	"strconv"
)

const (
	BaseMaxHP      = 20
	BaseMaxStamina = 50
)

type Player struct {
	name    string
	hp      float64
	stamina float64
	xp      int
}

func (p *Player) getInfo() string {
	info := ""
	info += p.name + "\n"
	info += "\thp:      " + strconv.FormatFloat(p.hp, 'f', 1, 64)
	info += "\n\tstamina: " + strconv.FormatFloat(p.stamina, 'f', 1, 64)
	info += "\n\txp:      " + strconv.Itoa(p.xp)
	return info
}

func MakePlayer(name string) *Player {
	player := &Player{name, BaseMaxHP, BaseMaxStamina, 0}
	return player
}
