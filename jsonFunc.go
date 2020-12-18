package main

import (
	"encoding/json"
	"fmt"
)

type Res struct {
	Code int
	Message string
}


func main() {
	var r1 Res
	r1.Code = 500
	r1.Message = "what"
	toJson(&r1)
	setJson(&r1)
	toJson(&r1)
}

func toJson(r1 *Res) {
	jsons, _ := json.Marshal(r1)
	fmt.Println(string(jsons))
}

func setJson(res *Res){
	res.Message = "code"
	res.Code = 500
	fmt.Println(res)
}