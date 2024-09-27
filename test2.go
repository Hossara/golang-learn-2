package main

func main() {
	is64bit := ^uint(0)>>63 == 1
	println(is64bit)
}
