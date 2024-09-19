package main

func main() {
	m1 := map[int]string{}

	m1[1] = "one"
	m1[1000] = "on one"

	m2 := map[string]string{
		"first_name": "Hossein",
		"last_name": "Araghi",
	}

	println(m2["first_name"])
	println(m2["last_name"])

	m3 := make(map[string]string)
	m3["salam"] = m2["first_name"]

	println("Salam", m3["salam"])

	delete(m2, "first_name")

	value, exists := m2["first_name"]

    if exists {
        println("First name is", value)
    } else {
		println("error")
	}

	m4 := make(map[string]map[string]int)

	m4["map1"] = map[string]int {
		"d1": 2,
	}
}