package main

import (
	"fmt"
	"math"
)

func calculateTax(p int) int {
	tax := 0

	if p > 100000 {
		tax += int(math.Floor(0.2 * float64(p-100000)))
		p = 100000
	}
	if p > 50000 {
		tax += int(math.Floor(0.15 * float64(p-50000)))
		p = 50000
	}
	if p > 10000 {
		tax += int(math.Floor(0.1 * float64(p-10000)))
		p = 10000
	}
	tax += int(math.Floor(0.05 * float64(p)))

	return tax
}

func main() {
	var p int
	fmt.Scan(&p)

	tax := calculateTax(p)
	fmt.Println(tax)
}