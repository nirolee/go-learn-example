package main

import "fmt"

func variable() {
	var a int = 3
	var b string = "test"
	fmt.Printf("%d %q", a, b)
}

func variableTypeDeduction() {
	a, b, c, s := 3, 5, true, "true"
	b = 4
	fmt.Println(a, b, c, s)
}

func grade(score int) string {
	grade := ""
	switch {
	case score < 60:
		grade = "f"
	case score < 70:
		grade = "b"
	case score < 100:
		grade = "a"
	default:
		panic(fmt.Sprintf("wrong score %d", score))
	}
	return grade
}
func main() {
	variable()
	variableTypeDeduction()
	grade(100)
}
