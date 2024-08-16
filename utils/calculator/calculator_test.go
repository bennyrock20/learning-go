package calculator

import "testing"

func TestAdd(t *testing.T) {
	calc := Calculator{}
	result := calc.Add(2, 3)
	expected := 5

	if result != expected {
		t.Errorf("Add(2, 3) = %d; want %d", result, expected)
	}
}
