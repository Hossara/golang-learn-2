package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	fmt.Fscanf(os.Stdin, "%s")
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Text()
}