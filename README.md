# math-package

This repository is just a test I made while learning how to publish a Go package.

## Downloading the package
```bash
go get -u github.com/Liedsonfsa/math-package@v0.1.1
```

## Package Structures

```Go
// Calculator represents the calculator
type Calculator struct {}
```

## Package functions

```Go
// NewCalculator create a new calculator
func NewCalculator() *Calculator {
    return &Calculator{}
}
```

## Package Methods

```Go
// SumInt performs the sum of the sequence of numbers
func (c Calculator) SumInt(numbers ...int) int {
    total := 0

	for _, n := range numbers {
		total += n
	}

	return total
}
```

## Using the package
```Go
package main

import (
    "fmt"

    "github.com/Liedsonfsa/math-package"
)

func main() {
    calculator := math.NewCalculator()

    sum := calculator.SumInt(1, 2, 3, 4, 5)

    fmt.Println(sum)
}
```