package main

func main() {
	// infinitie
    for {
        println("forever")
		break
    }
	
	i1 := 0
    for i1 <= 10 {
        println(i1, ")")
		i1++
    }
	
	println("--------")
	
	for i2 := 0; i2 <= 5; {
		println(i2, ")")
		i2++
	}
	
	println("--------")
	
	for i3 := 0; i3 <= 5; i3++{
		println(i3, ")")
		i3++
	}
}