package check

import (
	"bytes"
	"encoding/json"
	"io"
	"os"

	"krestNulls/game"
)

/*Проверка уровня на корректность*/
func CheckJSON(path string) bool {
	var (
		result game.GameMap //Результат
		buf    bytes.Buffer //Буффер
	)
	level, erOpenFile := os.Open(path) //Открываем json файл с уровнем
	decoder := json.NewDecoder(&buf)   //Иницилизируем декодер
	if erOpenFile != nil {
		return false //Неправильный адрес
	}
	defer level.Close()

	_, er := io.Copy(&buf, level) //Копируем из файла в буффер
	if er != nil {
		return false
	}

	err := decoder.Decode(&result) //Декодирум json из буффера в структуру GameMap
	if err != nil {
		return false
	}
	if len(result.GameMap) == 0 || len(result.WinCombinations) == 0 || result.Size <= 0 {
		return false
	}
	return true
}
