package main

import (
	"fmt"
	"strconv"
)

func convertBin(n int) string {
	result := ""
	for ; n > 0; n /= 2 {
		lsb := n % 2
		result = strconv.Itoa(lsb) + result
	}
	return result
}

func main() {
	fmt.Println(
		convertBin(5),
		convertBin(13),
	)
}
