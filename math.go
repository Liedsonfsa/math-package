package math

// Calculator represents the calculator
type Calculator struct {}

// NewCalculator create a new calculator
func NewCalculator() *Calculator {
	return &Calculator{}
}

// SumInt performs the sum of the sequence of numbers
func (c *Calculator) SumInt(numbers ...int) int {
	total := 0

	for _, n := range numbers {
		total += n
	}

	return total
}
