package math

const (
	PI float64 = 3.14159265359
)

type Calculator struct {}

func NewCalculator() *Calculator {
	return &Calculator{}
}

func (c *Calculator) SumInt(numbers ...int) int {
	total := 0

	for _, n := range numbers {
		total += n
	}

	return total
}
