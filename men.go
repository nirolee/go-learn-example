package main

import (
	"fmt"
	"runtime"
)

func printStat (men runtime.MemStats){
	runtime.ReadMemStats(&men)
	fmt.Println("mem.Alloc" , men.Alloc)
	fmt.Println("mem.HeapAlloc" , men.HeapAlloc)
	fmt.Println("mem.HeapIdle" , men.HeapIdle)
}


func main(){
	var mem runtime.MemStats
	for i:=0; i < 10; i++ {
		s := make([]byte, 100000000)
		if s == nil{
			fmt.Println("fail")
		}
		printStat(mem)
	}

}