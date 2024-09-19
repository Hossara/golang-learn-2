package main

import "fmt"

func main() {
	// Lenght should be pre defined
	var arr1 [10]int

	arr1[0] = 5
	arr1[1] = 10
	arr1[9] = 11

	fmt.Printf("%v\n", arr1)
	println(arr1[1])

    for i := 0; i <= 9; i++ {
		println(arr1[i])
    }

	println("------------------------------")

	arr2 := [...]int{1, 2, 3}

	for i := 0; i < len(arr2); i++ {
		println(arr2[i])
    }

	println("------------------------------")

	for i := range arr1 {
		println(i, arr1[i])
	}
}