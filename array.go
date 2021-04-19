package main

var (
	size     = 1024
	max_size = size * 2
)

type a struct {
	age  int
	name string
}

type b struct {
	age  int
	name string
}

type Int1 int
type Int2 = int

func main() {
	//s1 := []int{1, 2, 3}
	//s2 := []int{1, 2, 4}
	//s1 = append(s1, s2[:]...)
	//fmt.Println(size, max_size)
	//c := a{1, "1"}
	//d := a{1, "1"}
	//if c==d {
	//	println(1)
	//}
	//var i int = 2
	//var i2 Int1 = i
	//var i2 Int2 = i
	//fmt.Println(i2)

}
