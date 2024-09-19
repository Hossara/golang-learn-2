package main

import (
    "bufio"
    "fmt"
    "os"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)

	scanner.Scan()
	println(scanner.Text())
	scanner.Scan()
	println(scanner.Text())

	println("Enter a number")

	var i int
	fmt.Scanf("%d", &i)
	println(i)

	print("end")
}