package main

type myInt int64

func main() {
	// intSlice := []int{1, 2, 3, 4}
	// strSlice := []string{"mmd", "asqar", "ahmad"}
	// floatSlice := []float64{1.1, 2.2, 3.3}

	// printElements(intSlice)
	// printElements(strSlice)
	// printElements(floatSlice)

	// intSlice := []int{1, 2, 3, 4}
	// strSlice := Map[int, string](intSlice, func(i int) string {
	// 	return "converted -> " + strconv.Itoa(i)
	// })


	println(min(min(2, 3), min(3, 4)))
}

// Ype parametric programming --> generics
func printElements[T any](s []T) {
	for _, v := range s {
		println(v)
	}
}

func Map[T, U any](s []T, f func(T) U) []U {
	resault := make([]U, len(s))

	for i := range s {
		resault[i] = f(s[i])
	}

	return resault
}

// ~ = int or subtypes
// type Ints interface {
// 	~int | ~int8 | ~int16 | ~int32 | ~int64 
// }

type Numbers interface {
	~int | ~int8 | ~int16 | ~int32 | ~int64 |
	~uint | ~uint8 | ~uint16 | ~uint32 | ~uint64 |
	~float32 | ~float64
}

func min[T Numbers](a, b T) T {
	if a < b {
		return a
	}
	return b
}

// type Numbers interface {
// 	int | uint | float64
// }

// func min[T Numbers](a, b T) T {
// 	if a < b {
// 		return a
// 	}
// 	return b
// }

// func min[T int | uint | float64](a, b T) T {
// 	if a < b {
// 		return a
// 	}
// 	return b
// }
