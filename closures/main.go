package main

import "time"

// closure: Function which returned by a function
// closure inherits the parent function scope + its own scope

func main() {
	m := map[string]int{
		"Hossein": 18,
		"Narsis":  21,
		"Sahar":   23,
	}

	checkCache := cacheChecker(m)

	println(checkCache("Dara"))
	println(checkCache("Hossein"))

	seq := getSequence(0, 20, 3)

	for {
		v, ok := seq()

		if !ok {
			println(v)
			break
		}
		println(v)
	}

	// Decorator
	delay(func() {
		println("This is my action")
	})()
}

func cacheChecker(m map[string]int) func(key string) int {
	return func(key string) int {
		if v, ok := m[key]; ok {
			return v
		}
		return -1
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

func delay(f func()) func() {
	return func() {
		println("Starting...")
		time.Sleep(time.Second * 2)
		f()
		println("Done")
	}
}
