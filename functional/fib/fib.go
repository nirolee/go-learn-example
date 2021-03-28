package fib

import "fmt"

func Fibonacci() func() int {
	a, b := 0, 1
	return func() int {
		a, b = b, a+b
		fmt.Println(a, b)
		return a
	}
}
