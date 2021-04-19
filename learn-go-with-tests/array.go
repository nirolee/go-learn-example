package main

func Sum(nums []int) int {
	sum := 0
	for _, num := range nums {
		sum += num
	}
	return sum
}

func Sum1(nums ...[]int) (sum []int) {
	lnums := len(nums)
	s := make([]int, lnums)
	for i, num := range nums {
		s[i] = Sum(num)
	}
	return s
}
