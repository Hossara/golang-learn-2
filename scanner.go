package main

import (
    "bufio"
    "os"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	// scanner.Scan()

	// println(scanner.Text())

	var (
		firstName string
		lastName string
	)

	/*for print("Enter your fisrt name: "); scanner.Scan(); {
        if firstName == "" {
			print("Enter your fisrt name: ")
			firstName = scanner.Text()
        } else { break }
	}*/

	for {
        if firstName == "" {
			print("Enter your fisrt name: ")
			scanner.Scan()
			firstName = scanner.Text()
        } else { break }
	}

	for {
        if lastName == "" {
			print("Enter your last name: ")
			scanner.Scan()
			lastName = scanner.Text()
        } else { break }
	}

	println("Your fullname: ", firstName, lastName)
}