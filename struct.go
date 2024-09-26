package main

type Persion struct {
	FirstName string
	LastName string

	//X
}

type X struct {
     Y int
}

func main() {
	// Way 1
	p := Persion{
		FirstName: "Hossein",
		LastName: "Araghi",
	}
	println(p)

	// Way 2
	p2 := Persion{"Hossein", "Araghi"}
	println(p2)

	// Way 3
	var p3 Persion
	p3.FirstName = "Hossein"
	p3.LastName = "Araghi"

	// Way 5
	// Created persion zero value
	p4 := new(Persion) // p4 = &Persion{}
	println(p4)

	// Way 6
	// Created a pointer to initilize later with nil value
	var p5 *Persion
	println(p5)
}