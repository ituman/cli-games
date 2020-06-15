package seabattle

import (
	"fmt"
)

type Player interface {
	getEnemy() Player
	getField() *Field
	getFleet() Fleet
	getName() string
	initField()
	shot() bool
	isLose() bool
}

func Game() {
	bot := Bot{
		name:  "bot",
		field: NewField(),
		fleet: NewFleet(),
	}
	me := Me{
		name:  "Ivan",
		field: NewField(),
		fleet: NewFleet(),
	}

	bot.enemy = &me
	me.enemy = &bot

	bot.initField()
	me.initField()
	var thisPlayer Player

	thisPlayer = &bot

	for {
		fmt.Println("Ходит ", thisPlayer.getName())
		for {
			isHit := thisPlayer.shot()
			if !isHit {
				break
			}
		}
		if thisPlayer.getEnemy().isLose() {
			fmt.Println(thisPlayer.getEnemy().getName(), " - проиграл!")
			break
		}
		printGameField(me)
		thisPlayer = thisPlayer.getEnemy()
	}
}

func printGameField(me Me) {
	fmt.Println("Поле противника:")
	fmt.Println(me.getEnemy().getField().Print(false))
	fmt.Println("Твое поле:")
	fmt.Println(me.getField().Print(true))
	fmt.Println("Осталось кораблей у противника: ", me.getEnemy().getFleet().StillAlive())
}
