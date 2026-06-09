package main

import "fmt"

func Map[T any, U any](slice []T, action func(T) U) []U {
	res := make([]U, len(slice))
	for i, v := range slice {
		res[i] = action(v)
	}
	return res
}

func Curry[T any, U any, R any](f func(T, U) R) func(T) func(U) R {
	return func(t T) func(U) R {
		return func(u U) R {
			return f(t, u)
		}
	}
}

func Filter[T any](slice []T, predicate func(T) bool) []T {
	res := make([]T, 0)
	for _, v := range slice {
		if predicate(v) {
			res = append(res, v)
		}
	}
	return res
}
func Fold[T any, U any](slice []T, action func(U, T) U, initial U) U {
	result := initial
	for _, v := range slice {
		result = action(result, v)
	}
	return result
}

func FoldRight[T any, U any](slice []T, action func(T, U) U, initial U) U {
	result := initial
	for i := len(slice) - 1; i >= 0; i-- {
		result = action(slice[i], result)
	}
	return result
}

func Zip[T any, U any](a []T, b []U) [][2]any {
	length := len(a)
	if len(b) < length {
		length = len(b)
	}
	res := make([][2]any, length)
	for i := 0; i < length; i++ {
		res[i] = [2]any{a[i], b[i]}
	}
	return res
}

func main() {
	fmt.Printf("hello\n")
	xs := []int{1, 2, 3}
	f := Curry(func(x int, y int) int { return x + y })	
	ys := Map(xs, f(2))
	fmt.Printf("Result: %v\n", ys)

	xs2 := []int{1, 2, 3, 4}
	sum := Fold(xs2, func(acc int, x int) int { return acc + x }, 0)
	fmt.Printf("Sum: %d\n", sum)

	xs3 := []int{1, 2, 3, 4, 5}
	evens := Filter(xs3, func(x int) bool { return x%2 == 0 })
	fmt.Printf("Evens: %v\n", evens)

	fmt.Printf("Squared Evens: %v\n", Map(Filter(xs3, func(x int) bool { return x%2 == 0 }), func(x int) int { return x * x }))
	fmt.Printf("Sum of Squared Evens: %v\n", Fold(Map(Filter(xs3, func(x int) bool { return x%2 == 0 }), func(x int) int { return x * x }), func(acc int, x int) int { return acc + x }, 0))

	xs4 := []string {"a", "b", "c"}
	zipped := Zip(xs3, xs4)
	fmt.Printf("Zipped: %v\n", zipped)

	fmt.Printf("doubled: %v\n", Fold(xs3, func(acc []int, x int) []int { return append(acc, x*2) }, []int{}))

	xs5 := []string{"apple", "banana", "apple", "orange", "banana", "apple"}
	fmt.Printf("Unique: %v\n", Fold(xs5, func(acc []string, x string) []string {
		for _, v := range acc {
			if v == x {
				return acc
			}
		}
		return append(acc, x)
	}, []string{}))

}
