package main

import "fmt"


func ApplyMap[T any, U any](slice []T, action func(T) U) []U {
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


func main() {
	fmt.Printf("hello\n")
	xs := []int{1, 2, 3}
	f := Curry(func(x int, y int) int { return x + y })	
	ys := ApplyMap(xs, f(2))
	fmt.Printf("Result: %v\n", ys)
}
