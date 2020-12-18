package main

import (
	"fmt"
	"encoding/json"
)

func main() {
	//var j = map[string]interface{}{
	//	"username":"1j",
	//	"23":"2j",
	//}

	var j = map[int]string{
		1 : "第一个",
		2 : "第二个",
		3 : "第三个",
	}
	delete(j, 1)
	j[1] = "new 第一个"
	j[2] = "new 第2个"
	//res :=make(map[string]interface{})
	//res["code"] = 200
	//res["message"] = "success"
	//res["data"] = j
	jsons,_ := json.Marshal(j)
	fmt.Println(string(jsons))
}
