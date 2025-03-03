package math

import "testing"

func TestSumInt(t *testing.T) {
	sum := NewCalculator().SumInt(1, 2, 3, 4, 5)
	expected := 15

	if sum != expected {
		t.Errorf("expected '%d' but got '%d'", expected, sum)
	}
}