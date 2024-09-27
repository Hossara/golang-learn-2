package main

import (
	"fmt"
)

type OrderStatus uint8

// Enum = Contant Group + iota
const (
	OrderStatusAccepted OrderStatus = iota + 1
	OrderStatusPaied
)

func main() {
	s := []int{1, 2, 3, 4, 5, 6}

	for i, v := range s {
		fmt.Println("key: ", i, "value: ", v)
	}

	println("------------------")

	m := map[string]int{
		"Hossein": 18,
		"Narsis":  21,
		"Sahar":   23,
	}

	for k, v := range m {
		fmt.Println("key: ", k, "value: ", v)
	}

	println("------------------")

	for i := range 20 {
		println("Num: ", i)
	}

	println("------------------")

	seq := getSequence(5, 25, 5)
	for i, ok := seq(); ok && i < 21; i, ok = seq() {
		println(i)
	}

	println("------------------")

	var status OrderStatus

	switch status {
	case OrderStatusAccepted:
		println("Accepted")
		fallthrough // Run also next condition - like you dont write break
	case OrderStatusPaied:
		println("Paid")
	default:
		println("Canceled")
	}
}

func getSequence(start, stop, step int) func() (int, bool) {
	return func() (int, bool) {
		if start > stop {
			return stop, false
		}

		start += step

		if start > stop {
			return stop, false
		}

		return start, true
	}
}
