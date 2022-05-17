package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
	"strconv"
)

const (
	BaseMaxHP        = 100
	BaseMaxStamina   = 50
	BaseDamage       = 10
	BaseStaminaRegen = 3
	BaseLvl          = 1
	SaveFilePostFix  = ".kds"
	MySecret         = "abc&1*~#^2^#s0^=)^^7%b34"
)

type Player struct {
	Name         string
	Hp           float64
	Stamina      float64
	StaminaRegen float64
	Damage       float64
	Xp           int
	Lvl          int
}

func (p *Player) getInfo() string {
	info := ""
	info += p.Name + "\n"
	info += "\thp:            " + strconv.FormatFloat(p.Hp, 'f', 1, 64)
	info += "\n\tstamina:       " + strconv.FormatFloat(p.Stamina, 'f', 1, 64)
	info += "\n\tstamina regen: " + strconv.FormatFloat(p.StaminaRegen, 'f', 1, 64)
	info += "\n\tdamage:        " + strconv.FormatFloat(p.Damage, 'f', 1, 64)
	info += "\n\txp:            " + strconv.Itoa(p.Xp)
	info += "\n\tlvl:           " + strconv.Itoa(p.Xp)
	return info
}

func (p *Player) saveToFile(savefileName string) error {
	file, _ := json.MarshalIndent(*p, "", " ")
	encryptedFile, _ := Encrypt(string(file), MySecret)
	err := ioutil.WriteFile(savefileName, []byte(encryptedFile), 0644)

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

func loadPlayerFromFile(savefileName string) (*Player, error) {
	file, _ := ioutil.ReadFile(savefileName)
	decryptedFile, _ := Decrypt(string(file), MySecret)
	player := &Player{}
	err := json.Unmarshal([]byte(decryptedFile), player)
	return player, err
}

// TODO: Add error handling
func MakePlayer(name string) *Player {
	savefileName := name + SaveFilePostFix
	var player *Player
	if doesFileExist(savefileName) {
		fmt.Println("Reading from file...")
		player, _ = loadPlayerFromFile(savefileName)
	} else {
		player = &Player{name, BaseMaxHP, BaseMaxStamina, BaseStaminaRegen, BaseDamage, 0, BaseLvl}
		player.saveToFile(savefileName)
	}

	return player
}
