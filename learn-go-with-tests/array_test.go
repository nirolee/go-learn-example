package main

import (
	"reflect"
	"testing"
)

func TestSum(t *testing.T) {
	arr := []int{1, 2, 3, 4, 5}
	res := Sum(arr)
	expect := 15
	if res != expect {
		t.Errorf("arr %v, res %d", arr, res)
	}
}

func TestSum2(t *testing.T) {
	arr := []int{1, 2, 3, 4, 5}
	arr1 := []int{6, 7, 8}
	res := Sum1(arr, arr1)
	expect := []int{15, 21}
	if !reflect.DeepEqual(res, expect) {
		t.Errorf("res %v, expect %v", res, expect)
	}

}
