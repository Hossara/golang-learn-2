package main

type Persion2 struct {
	FistName string
	LastName string
	Age      uint
}

func (person Persion2) fullName() string {
	return person.FistName + " " + person.LastName
}

func (persion Persion2) getMamad() string {
	return "HI mmd, " + persion.FistName
}

func (person *Persion2) setAge(age int) {
	if age < 0 {
		age = 7
	}

	person.Age = uint(age)
}

func main() {
	person := new(Persion2)
	person.FistName = "Hossein"
	person.LastName = "Araghi"

	println(person.fullName())
	println(person.getMamad())

	println(person.Age)

	person.setAge(23)
	println(person.Age)
}
