package main

import (
	"encoding/json"
	"fmt"
)

type Data struct {
	Code int `json:"code"`
	Message string `json:"message"`
}

func main() {
	var reg Data
	reg.Code = 200
	reg.Message = "正常"

	a, err :=json.Marshal(reg)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(string(a))
	var b Data
	err = json.Unmarshal(a, &b)
	if err != nil{
		fmt.Println(err)
	}
	fmt.Println(b)
}
