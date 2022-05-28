package main

import (
	"fmt"
	"math/rand"
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

	GetName() string
	GetHp() float64
	GetLvl() int
	GetDamage() float64

	SetHp(float64)
	SetName(string)
	SetLvL(int)
	SetDamage(float64)
}

func InitEnemyRegistery() {
	myTypes := []interface{}{BackstreetFighter{}}
	for _, v := range myTypes {
		EnemyRegistery[reflect.TypeOf(v).Name()] = reflect.TypeOf(v)
	}
}

func GetRandomEnemy() Enemy {
	k := rand.Intn(len(EnemyRegistery))
	for key := range EnemyRegistery {
		if k == 0 {
			return MakeEnemy(key)
		}
		k--
	}
	panic("unreachable")
}

func MakeEnemy(enemyName string) Enemy {
	fmt.Println(EnemyRegistery[enemyName])
	v := reflect.New(EnemyRegistery[enemyName]).Elem().Addr()
	enemy := v.Interface().(Enemy)
	enemy.SetName(strings.Join(camelcase.Split(enemyName)[:], " "))
	enemy.SetLvL(10)

	return enemy
}
