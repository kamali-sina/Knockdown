package main

import (
	"fmt"
	"reflect"
	"strings"

	"github.com/fatih/camelcase"
)

const (
	BackstreetFighterName = "Backstreet Fighter"
)

var EnemyRegistery = make(map[string]reflect.Type)

type Enemy interface {
	ChooseMove()
}

type BackstreetFighter struct {
	Name   string
	Hp     float64
	Damage float64
	Lvl    int
}

func (bf BackstreetFighter) ChooseMove() {
	fmt.Println("choosing move")
}

func InitEnemyRegistery() {
	myTypes := []interface{}{BackstreetFighter{}}
	for _, v := range myTypes {
		EnemyRegistery[reflect.TypeOf(v).Name()] = reflect.TypeOf(v)
	}
}

func MakeEnemy(enemyName string) Enemy {
	v := reflect.New(EnemyRegistery[enemyName]).Elem()
	v.FieldByName("Name").SetString(strings.Join(camelcase.Split(enemyName)[:], " "))
	v.FieldByName("Lvl").SetInt(10)

	enemy := v.Interface().(Enemy)

	return enemy
}
