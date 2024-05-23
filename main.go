// go run main.go
package main

import (
	"bufio"
	"fmt"
	check "krestNulls/checkLevel"
	"krestNulls/game" //Пакет с игровыми функциями
	"os"
	"strconv"
)

func main() {
	var (
		gameMap   game.GameMap //Карта поля
		shoot     int          //В какую ячейку выстрел
		nameOfMap string       //Имя для игровой карты
		gameType  string       //Тип действия
	)
	sc := bufio.NewScanner(os.Stdin)                   //Сканер из стандартного ввода
	fmt.Println("ДОБРО ПОЖАЛОВАТЬ В КРЕСТИКИ НОЛИКИ!") //Приветствие

	for { //Начало бесконечного цикла
		for { //Ждём адекватный уровень
			fmt.Print("Уровень: ") //Какой уровень
			sc.Scan()
			nameOfMap = sc.Text()
			if nameOfMap == "f" { //Если f, то уровень с компьютера
				var (
					f    string //Имя уровня
					errr error  //Ошибка
				)
				sc.Scan()
				f = sc.Text()
				gameMap, errr = game.GetUserLevel(f + ".json") //Берём пользовательский уровень
				if errr != nil {
					fmt.Println(errr)
				} else {
					fmt.Printf("Уровень %v успешно загружен!\n", f)
					gameType = "game"
					break
				}
			} else if nameOfMap == "check" { //Для проверки корректности уровня
				gameType = "check"
				break
			} else {
				_, ok := game.Levels[nameOfMap]
				if ok {
					gameMap = game.Levels[nameOfMap]
					for i := range gameMap.GameMap { //Заполняем карту нулями
						gameMap.GameMap[i] = 0
					}
					gameType = "game" //Устанавливем тип игры как "game"
					break
				} else {
					fmt.Printf("Уровня %v ещё нет((\n", nameOfMap)
				}
			}
		}

		switch gameType { //Какой тип игры
		case "game": //Если это игра
			var i int = 1 //На случай если компьютер первый
			for {         //Бесконечный цикл для ставки ноликов и крестиков
				if gameMap.First && i == 1 {
					gameMap.ComputerDoShoot()
					i-- //Чтобы больше компьютер не ходил в начале хода
				}
				fmt.Println("Твой ход!")
				fmt.Println(gameMap.PrintMap()) //Печатаем карту первый раз
				for {                           //Делаем бесконечный цикл и ждём пока пользователь не введёт что-то адекватное
					fmt.Print("Введи куда ставить нолик: ")
					sc.Scan()
					shoot64, er := strconv.ParseInt(sc.Text(), 10, 64)
					if er == nil {
						shoot = int(shoot64)
					}
					if er != nil {
						fmt.Println("Неправильная координата стрельбы")
					} else if shoot == -1 { //Прерываем цикл если хотим выйти из партии
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
		case "check": //Проверка корректности уровня
			fmt.Print("Введи название уровня для проверки: ")
			sc.Scan()
			name := sc.Text()
			res := check.CheckJSON(name + ".json")
			if res {
				fmt.Printf("Уровень %v корректный\n", name)
			} else {
				fmt.Printf("Уровень %v не корректный\n", name)
			}
		}

	}
}
