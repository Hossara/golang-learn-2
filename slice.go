package main

func main() {
	slice := []int{1, 2, 3, 4, 5, 6, 7, 8}
    slice = append(slice, 5)

	for i := 0; i < len(slice); i++ {
		println(slice[i])
	}
}