package game

import (
	"crypto/rand"
	"fmt"
	"math/big"
)

/*Структура для игрового поля*/
type GameMap struct {
	GameMap         []int   `json:"map"`          //Карта игроков, 0 - пусто, 1 - человек, 2 - компьютер
	WinCombinations [][]int `json:"combinations"` //Победные комбинации
	Size            int     `json:"size"`         //Размер поля в длину
}

/*Функция печатает карту для пользователя*/
func (m GameMap) PrintMap() string {
	var res string //Результат
	for ind, val := range m.GameMap {
		if ind%m.Size == 0 && ind != 0 { //Проверка на чётность размеру поля
			res = fmt.Sprintf("%s\n", res)
		}
		if val == 0 {
			res = fmt.Sprintf("%s%v ", res, ind+1)
		} else if val == 1 {
			res = fmt.Sprintf("%sO ", res)
		} else if val == 2 {
			res = fmt.Sprintf("%sX ", res)
		}
	}
	return res
}

/*
Функция определяет результат игры
0 - ничего не произошло
1 - победа
2 - ничья

Входные параметры
1 - человек
2 - компьютер
*/
func (m GameMap) Check(why int) int {
	var sum int //Для определения ничьи

	for _, combination := range m.WinCombinations { //Перебор комбинаций
		for j, val := range combination { //Перебор внутреннего среза
			if m.GameMap[val] != why { //Если клетка не соответствует чему надо, то выходим
				break
			}
			if j == len(combination)-1 { //Если все клетки заполнены чем нужно
				return 1
			}
		}
	}
	for _, val := range m.GameMap { //Перебор карты поля на ничью
		if val != 0 { //Если поле не свободно, то увеличиваем счётчик
			sum++
		}
	}
	if sum == len(m.GameMap) { //Если заполнены все клетки
		return 2
	}
	return 0 //По умолчанию ничего не произошло
}

/*Функция ставит O или X куда скажет пользователь/компьютер*/
func (m *GameMap) Put(why, where int) {
	m.GameMap[where] = why
}

/*Функция делает ход за компьютер*/
func (m GameMap) ComputerDoShoot() {
	/*Сначала проверяем есть ли варианты для победы*/
	var success bool = m.searchOption(2)
	if success {
		return
	}
	/*Мешаем выиграть человеку*/
	success = m.searchOption(1)
	if success {
		return
	}
	/*Иначе бьём в случайное место*/
	for {
		nBig, err := rand.Int(rand.Reader, big.NewInt(int64(len(m.GameMap)-1))) //Генерируем случайное число
		if err != nil {
			panic(err)
		}
		whereShoot := int(nBig.Int64())
		if m.GameMap[whereShoot] == 0 { //Если поле свободно, то бьём туда и выходим из цикла
			m.Put(2, whereShoot)
			break
		}
	}
}

/*Для выигрыша компьютера / мешаем человеку выиграть*/
func (m GameMap) searchOption(why int) bool {
	for _, option := range m.WinCombinations {
		var null int = -1 //Индекс первой свободной клетки
		var sum int       //Сумма свободных клеток в комбинации
		for _, val := range option {
			if m.GameMap[val] == why { //Если там человек/компьютер, увеличиваем счётчик
				sum++
			}
			if m.GameMap[val] == 0 { //Нашли свободную клетку
				null = val
			}
		}
		if null != -1 && sum == len(option)-1 { //Если есть 1 свободная клетка и все остальные заняты, то бьём в свободную
			m.Put(2, null)
			return true
		}
	}
	return false
}
