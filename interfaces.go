package main

type Human interface {
	Name() string
	setAge(age int)
	SayHello() string
	PrintAge() uint
}

type Persion3 struct {
	FistName string
	LastName string
	Age      uint
}

type AsianHuman interface {
	Human
	Language() string
}

type IranianPerson struct {
	Persion3
	language string
}

func (person IranianPerson) Language() string {
	return "Persian"
}

func (person IranianPerson) SayHello() string {
	return "Hello"
}

func (person IranianPerson) setAge(age int) {
	if age < 0 {
		age = 7
	}

	person.Age = uint(age)
}

func (person Persion3) Name() string {
	return person.FistName + " " + person.LastName
}

func (persion Persion3) SayHello() string {
	return "HI mmd, " + persion.FistName
}

func (person *Persion3) setAge(age int) {
	if age < 0 {
		age = 7
	}

	person.Age = uint(age)
}

func (person Persion3) PrintAge() uint {
	return person.Age
}

func main() {

	var h Human // nil

	//h = &Persion3{}
	h = IranianPerson{}

	h.setAge(23)

	println(h.PrintAge())

	var h2 AsianHuman

	h2 = IranianPerson{}

}
