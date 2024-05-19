package game

import "testing"

func TestPut(t *testing.T) {
	for j, m := range Levels {
		t.Run(j, func(t *testing.T) {
			m.Put(1, 3)
			if m.GameMap[3] != 1 {
				t.Errorf("On test %v error: 3 not 1", j)
			}
		})
	}
	for j, m := range Levels {
		t.Run(j, func(t *testing.T) {
			m.Put(1, 2)
			if m.GameMap[2] != 1 {
				t.Errorf("On test %v error: 2 not 1", j)
			}
		})
	}
	for j, m := range Levels {
		t.Run(j, func(t *testing.T) {
			m.Put(1, 2)
			m.Put(2, 2)
			if m.GameMap[2] != 1 {
				t.Errorf("On test %v error: 2 not 1", j)
			}
		})
	}
}

func TestPrintMap(t *testing.T) {
	options := map[string]GameMap{
		"print empty 2x2": {
			GameMap:         []int{0, 0, 0, 0},
			WinCombinations: [][]int{{0, 1}, {2, 3}, {1, 2}, {0, 3}},
			Size:            2,
		},
		"print not empty 2x2": {
			GameMap:         []int{0, 1, 2, 1},
			WinCombinations: [][]int{{0, 1}, {2, 3}, {1, 2}, {0, 3}},
			Size:            2,
		},
		"print empty 3x3": {
			GameMap:         []int{0, 0, 0, 0, 0, 0, 0, 0, 0},
			WinCombinations: [][]int{{0, 1, 2}, {3, 4, 5}, {6, 7, 8}, {0, 3, 6}, {1, 4, 7}, {2, 5, 8}, {0, 4, 8}, {2, 4, 6}},
			Size:            3,
		},
		"print not empty 3x3": {
			GameMap:         []int{2, 1, 0, 0, 2, 1, 1, 0, 2},
			WinCombinations: [][]int{{0, 1, 2}, {3, 4, 5}, {6, 7, 8}, {0, 3, 6}, {1, 4, 7}, {2, 5, 8}, {0, 4, 8}, {2, 4, 6}},
			Size:            3,
		},
	}
	optionsResults := map[string]string{
		"print empty 2x2":     "1 2 \n3 4 ",
		"print not empty 2x2": "1 O \nX O ",
		"print empty 3x3":     "1 2 3 \n4 5 6 \n7 8 9 ",
		"print not empty 3x3": "X O 3 \n4 X O \nO 8 X ",
	}
	for y, test := range options {
		t.Run(y, func(t *testing.T) {
			res := test.PrintMap()
			if res != optionsResults[y] {
				t.Errorf("failded %v", y)
			}
		})
	}
}
