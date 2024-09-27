package main

type Person struct {
	FirstNama  string
}

func changeName(p *Person) {
	p.FirstNama = "mmd"
}

func main() {
	//p := new(Person) // p := &Persion

	p := Person{
		FirstNama:  "Hossein ",
	}

	p.FirstNama = "Hossein"
	println(p.FirstNama)

	changeName(&p)

	println(p.FirstNama)
}