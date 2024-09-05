package main

import (
	"fmt"
    "strings"
	"errors"
)

// Global variable
var (
	str1 string
	str2 = "hi"
	number int
	number2  = 12
)

func test1() (string, error) {
	return "nil", errors.New("hi, this is error")
}

func main() {
	// All type of string combining 
	println("Hello", "World")
	fmt.Println("Hello", "World")
	fmt.Printf("%v %v\n", "Hello", "World")
	
	x := fmt.Sprintf("%v %v", "Hello", "World")
	y := fmt.Sprintln("Hello", "World")
	println(x)
	print(y)
	
	println(strings.Contains(y, "Hello"))
	println(strings.TrimSpace(" dddd "))
	println(strings.Trim(" dddd --", "-"))
	println(strings.Fields(y)[0])

    if strings.Contains(y, "hello") {
        println("Contains")
    } else {
		println("Not contains")
	}

    if myss := "hello world"; strings.Contains(myss, "world") {
        println(myss, "This works")
    }
	// println(myss) There is no myss here
	
	if ss, err := test1(); err != nil {
		println("There is error when asking for", ss)
	}
}