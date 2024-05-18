// go run main.go gameMap.go levels.go
package main

import (
	"fmt"
)

func main() {
	var (
		gameMap   GameMap //Карта поля
		shoot     int     //В какую ячейку выстрел
		nameOfMap string  //Имя для игровой карты
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
		gameMap, errr = getUserLevel(f + ".json")
		if errr != nil {
			panic(errr)
		}
	} else {
		gameMap = levels[nameOfMap]
	}

	for {
		fmt.Println("Твой ход!")
		fmt.Println(gameMap.printMap())
		fmt.Print("Введи куда ставить нолик: ")
		fmt.Scan(&shoot)

		gameMap.put(1, shoot-1)

		if gameMap.check(1) == 1 {
			fmt.Println("Победа!")
			fmt.Println(gameMap.printMap())
			break
		} else if gameMap.check(1) == 2 {
			fmt.Println("Ничья!")
			fmt.Println(gameMap.printMap())
			break
		} else {
			gameMap.computerDoShoot()
			if gameMap.check(2) == 1 {
				fmt.Println("Поражение!")
				fmt.Println(gameMap.printMap())
				break
			} else if gameMap.check(2) == 2 {
				fmt.Println("Ничья!")
				fmt.Println(gameMap.printMap())
				break
			}
		}
	}
	var t int
	fmt.Scan(&t)
}
