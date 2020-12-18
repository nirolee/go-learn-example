package main

func main() {
	map_arr := map[int] string{
		1 : "第一个",
		2 : "第二个",
		3 : "第三个",
	}

	true_arr := [4] string {"one", "twe", "trues", "four"}

	//类似while

	for  {
		println(1)
		break
	}
	
	for i,v := range map_arr{
		println(i ,"+", v)
	}

	for i,v := range true_arr{
		println(i ,"+", v)
	}

	for i:=0 ; i<4 ; i++{
		if i == 2{
			goto END
		}
		println(true_arr[i])
	}

	END:
		println("this is end")
}
