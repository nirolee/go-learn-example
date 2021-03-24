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

func z() {
	var a = 2
	var pa *int = &a
	*pa = 3
	fmt.Println(a)
}
func swap(a, b int) (int, int) {
	return b, a
}
func main() {

	a, b := 1, 4
	a, b = swap(a, b)
	fmt.Println(a, b)
	//eval(12, 5, "/")
	//if result, err := eval(12, 5, "x"); err != nil {
	//	fmt.Println("error: ",  err)
	//} else {
	//	fmt.Println(result)
	//}
	//q, r := div(12, 5)
	//fmt.Println(q, r)
}
