package main

import "fmt"

type BackstreetFighter struct {
	Name   string
	Hp     float64
	Damage float64
	Lvl    int
}

func (bf *BackstreetFighter) ChooseMove() {
	fmt.Println("choosing backstreet move")
}

func (bf *BackstreetFighter) GetName() string {
	return bf.Name
}

func (bf *BackstreetFighter) GetHp() float64 {
	return bf.Hp
}

func (bf *BackstreetFighter) GetLvl() int {
	return bf.Lvl
}

func (bf *BackstreetFighter) GetDamage() float64 {
	return bf.Damage
}

func (bf *BackstreetFighter) SetHp(hp float64) {
	bf.Hp = hp
}

func (bf *BackstreetFighter) SetName(name string) {
	bf.Name = name
}

func (bf *BackstreetFighter) SetLvL(lvl int) {
	bf.Lvl = lvl
}

func (bf *BackstreetFighter) SetDamage(damage float64) {
	bf.Damage = damage
}
