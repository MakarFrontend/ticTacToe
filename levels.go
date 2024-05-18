package main

import (
	"bytes"
	"encoding/json"
	"errors"
	"io"
	"os"
)

/*Map для встроенных уровней*/
var levels map[string]GameMap = map[string]GameMap{
	"2x2": { //Поле 2x2
		GameMap:         []int{0, 0, 0, 0},
		WinCombinations: [][]int{{0, 1}, {2, 3}, {1, 2}, {0, 3}},
		Size:            2,
	},
	"3x3": { //Поле 3x3
		GameMap:         []int{0, 0, 0, 0, 0, 0, 0, 0, 0},
		WinCombinations: [][]int{{0, 1, 2}, {3, 4, 5}, {6, 7, 8}, {0, 3, 6}, {1, 4, 7}, {2, 5, 8}, {0, 4, 8}, {2, 4, 6}},
		Size:            3,
	},
}

/*Берём пользовательский уровень, возвращает структуру GameMapи ошибку*/
func getUserLevel(path string) (GameMap, error) {
	var (
		result GameMap      //Результат
		buf    bytes.Buffer //Буффер
	)
	level, erOpenFile := os.Open(path) //Открываем json файл с уровнями
	decoder := json.NewDecoder(&buf)   //Иницилизируем декодер
	if erOpenFile != nil {
		return GameMap{}, errors.New("Ошибка в поиске файла!") //Неправильный адрес
	}
	defer level.Close()

	_, er := io.Copy(&buf, level) //Копируем из файла в буффер
	if er != nil {
		return GameMap{}, errors.New("Неккоректный файл!")
	}

	decoder.Decode(&result) //Декодирум json из буффера в структуру GameMap
	return result, nil
}
