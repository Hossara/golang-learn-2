package main

import "fmt"

func main() {
	slice1 := make([]string, 10)

	slice1 = append(slice1, "hi")

	slice2 := make([]float64, 3) // len = 3 cap = 3
	slice3 := append(slice2, 9) // len = 4 cap = 6
	slice4 := append(slice3, 8) // len = 5 cap = 6
	slice4[0] = 5
	fmt.Println(slice2)
	fmt.Println(slice3)
	fmt.Println(slice4)
	println(cap(slice2))
	println(cap(slice3))
	println(cap(slice4))
	// Every slice is save in a memory called cap or capacity
	// When you append to a slice and the secound slice capacity changed, then you have two slices.
 }