package main

import "fmt"

func eval(a, b int, op string) (int, error) {
	switch op {
	case "+":
		return a + b, nil
	case "-":
		return a - b, nil
	case "*":
		return a * b, nil
	case "/":
		return a / b, nil
	default:
		return 0, fmt.Errorf("unsupported operator" + op)
	}
}

func div(a, b int) (q, r int) {
	return a / b, a % b
}

func main() {
	eval(12, 5, "/")
	if result, err := eval(12, 5, "x"); err != nil {
		fmt.Println("error: ", err)
	} else {
		fmt.Println(result)
	}
	q, r := div(12, 5)
	fmt.Println(q, r)
}
