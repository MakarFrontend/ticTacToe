// go run main.go
package main

import (
	"fmt"
	"krestNulls/game" //Пакет с игровыми функциями
	"os"
)

func main() {
	var (
		gameMap   game.GameMap //Карта поля
		shoot     int          //В какую ячейку выстрел
		nameOfMap string       //Имя для игровой карты
	)
	fmt.Println("ДОБРО ПОЖАЛОВАТЬ В КРЕСТИКИ НОЛИКИ!") //Приветствие

	for { //Начало бесконечного цикла
		for { //Ждём адекватный уровень
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
					fmt.Println(errr)
				} else {
					fmt.Printf("Уровень %v успешно загружен!\n", f)
					break
				}
			} else {
				_, ok := game.Levels[nameOfMap]
				if ok {
					gameMap = game.Levels[nameOfMap]
					for i := range gameMap.GameMap {
						gameMap.GameMap[i] = 0
					}
					break
				} else {
					fmt.Printf("Уровня %v ещё нет((\n", nameOfMap)
				}
			}
		}
		var i int = 1
		for { //Бесконечный цикл для ставки ноликов и крестиков
			if gameMap.First && i == 1 {
				gameMap.ComputerDoShoot()
				i--
			}
			fmt.Println("Твой ход!")
			fmt.Println(gameMap.PrintMap()) //Печатаем карту первый раз
			for {                           //Делаем бесконечный цикл и ждём пока пользователь не введёт что-то адекватное
				fmt.Print("Введи куда ставить нолик: ")
				fmt.Scan(&shoot)
				if shoot == -1 { //Прерываем цикл если хотим выйти из партии
					break
				} else if shoot == -2 { //-2 - свайп хода
					fmt.Fprintln(os.Stdout, "Ход пропущен")
					break
				} else if shoot > 0 && shoot <= len(gameMap.GameMap) {
					if gameMap.GameMap[shoot-1] == 0 {
						gameMap.Put(1, shoot-1) //Ставим нолик
						break
					} else {
						fmt.Println("Неправильная координата стрельбы")
					}
				} else {
					fmt.Println("Неправильная координата стрельбы")
				}
			}
			if shoot == -1 { //Прерываем партию если координата стрельбы -1
				break
			}

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
