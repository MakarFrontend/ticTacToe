// go run main.go
package main

import (
	"fmt"
	"krestNulls/game"
)

func main() {
	var (
		gameMap   game.GameMap //Карта поля
		shoot     int          //В какую ячейку выстрел
		nameOfMap string       //Имя для игровой карты
	)
	fmt.Println("ДОБРО ПОЖАЛОВАТЬ В КРЕСТИКИ НОЛИКИ!")

	fmt.Print("Уровень: ")
	fmt.Scan(&nameOfMap)
	if nameOfMap == "f" {
		var (
			f    string
			errr error
		)
		fmt.Scan(&f)
		gameMap, errr = game.GetUserLevel(f + ".json")
		if errr != nil {
			panic(errr)
		}
	} else {
		gameMap = game.Levels[nameOfMap]
	}

	for {
		fmt.Println("Твой ход!")
		fmt.Println(gameMap.PrintMap())
		fmt.Print("Введи куда ставить нолик: ")
		fmt.Scan(&shoot)

		gameMap.Put(1, shoot-1)

		if gameMap.Check(1) == 1 {
			fmt.Println("Победа!")
			fmt.Println(gameMap.PrintMap())
			break
		} else if gameMap.Check(1) == 2 {
			fmt.Println("Ничья!")
			fmt.Println(gameMap.PrintMap())
			break
		} else {
			gameMap.ComputerDoShoot()
			if gameMap.Check(2) == 1 {
				fmt.Println("Поражение!")
				fmt.Println(gameMap.PrintMap())
				break
			} else if gameMap.Check(2) == 2 {
				fmt.Println("Ничья!")
				fmt.Println(gameMap.PrintMap())
				break
			}
		}
	}
	var t int
	fmt.Scan(&t)
}
