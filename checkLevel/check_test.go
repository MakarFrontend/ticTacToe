package check

import "testing"

func TestCheckJSON(t *testing.T) {
	var optinos map[string]bool = map[string]bool{
		"test-false.json":  false,
		"test-true.json":   true,
		"test2-false.json": false,
		"test2-true.json":  true,
	}
	for key, val := range optinos {
		t.Run(key, func(t *testing.T) {
			res := CheckJSON(key)
			if res != val {
				t.Errorf("Ошибка в %v", key)
			}
		})
	}
}
