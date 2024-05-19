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
	fmt.Println("ДОБРО ПОЖАЛОВАТЬ В КРЕСТИКИ НОЛИКИ!") //Приветствие

	for { //Начало бесконечного цикла
		fmt.Print("Уровень: ") //Какой уровень
		fmt.Scan(&nameOfMap)
		if nameOfMap == "f" { //Если f, то уровень с компьютера
			var (
				f    string //Имя уровня
				errr error  //Ошибка
			)
			fmt.Scan(&f)
			gameMap, errr = game.GetUserLevel(f + ".json") //Берём пользовательский уровень
			if errr != nil {
				panic(errr)
			}
		} else {
			gameMap = game.Levels[nameOfMap] //Если не f то берём уровень из стандартных
		}

		for {
			fmt.Println("Твой ход!")
			fmt.Println(gameMap.PrintMap()) //Печатаем карту 1 раз
			fmt.Print("Введи куда ставить нолик: ")
			fmt.Scan(&shoot)
			if shoot == -1 { //Прерываем партию если координа стрельбы -1
				break
			}
			gameMap.Put(1, shoot-1) //Ставим нолик

			if gameMap.Check(1) == 1 { //Если победа
				fmt.Println("Победа!")
				fmt.Println(gameMap.PrintMap())
				break
			} else if gameMap.Check(1) == 2 { //Если ничья
				fmt.Println("Ничья!")
				fmt.Println(gameMap.PrintMap())
				break
			} else { //Если не победа или ничья
				gameMap.ComputerDoShoot()  //Компьютер стреляет
				if gameMap.Check(2) == 1 { //Поражение
					fmt.Println("Поражение!")
					fmt.Println(gameMap.PrintMap())
					break
				} else if gameMap.Check(2) == 2 { //Ничья
					fmt.Println("Ничья!")
					fmt.Println(gameMap.PrintMap())
					break
				}
			}
		}
	}
}
