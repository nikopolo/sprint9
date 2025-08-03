package main

import (
	"testing"
)

// Пишите тесты в этом файле
func TestGenerateRandomElements(t *testing.T) {
	result := generateRandomElements(0)
	if result != nil {
		t.Errorf("expected nil\n")
	}

	result = generateRandomElements(-1)
	if result != nil {
		t.Errorf("expected nil")
	}

	result = generateRandomElements(1000)
	if result == nil {
		t.Errorf("expected slice with %d elements", 1000)
	}
}

func TestMaximum(t *testing.T) {
	//emptyList := []int{}
	result := maximum([]int{})
	if result != 0 {
		t.Error("expected 0")
	}

	result = maximum([]int{100})
	if result != 100 {
		t.Error("expected 100")
	}

	result = maximum([]int{1, 5, 10})
	if result != 10 {
		t.Error("expected 10")
	}

	result = maximum([]int{5, 5, 5})
	if result != 5 {
		t.Error("expected 5")
	}
}
