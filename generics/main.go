package main

import "strconv"

// type myInt int64
type Account[T Numbers] struct {
	Balance T
	Owner   string
	// T --> Not allowed to embed
}

type FloatAccount = Account[float64]

type Content interface {
	~string | ~[]byte
}

type Reader[T Content] interface {
	Read(content T) error
}

func main() {
	intSlice := []int{1, 2, 3, 4}
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

	println(Reduce(intSlice, "", func (s int, result string) string  {
		return result + " " + strconv.Itoa(s)  
	}))
}

// Ype parametric programming --> generics
// func printElements[T any](s []T) {
// 	for _, v := range s {
// 		println(v)
// 	}
// }

// Instantiate generic types
var MapInts = Map[int, string]

func Map[T, U any](s []T, f func(T) U) []U {
	result := make([]U, len(s))

	for i := range s {
		result[i] = f(s[i])
	}

	return result
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


type reducerFunc[T, U any] func(currentElement T, currentResult U) U

func Reduce[T, U any](s []T, init U, f reducerFunc[T, U]) U {
	for i := range s {
		init = f(s[i], init)
	}

	return init
}