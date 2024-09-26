package main

import (
    "fmt"
)

func main() {
	var (
		hopeNumber int // Hope - 2 <= p <= 100
		maximumNumber int // If reach game must end - 2 <= q <= 100000
	) // When reach the p, whe shoud (p / x) times say hope

    _, err := fmt.Scanf("%d %d", &hopeNumber, &maximumNumber)
    if err != nil {
        panic("Format error, correct format: number number")
    }

	for i := 1; i <= maximumNumber; i++ {
        if i % hopeNumber == 0 {
            y := 1
			for y <= i / hopeNumber {
				fmt.Print("Hope ")
				y++
            }
			fmt.Println()
        } else {
			fmt.Println(i)
		}
	}
}