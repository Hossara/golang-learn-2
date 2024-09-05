package main

func main() {
	i := 5

	// No need to break
    switch i + 1 {
   		case 1:
        	println("one")
   		case 2:
        	println("two")
        	println("again")
   		case 6:
        	println("six")
	    default:
	        println("Unknown number")
    }
}