package main

import (
	"fmt"
	"net/http"
	"net/http/httputil"
)

func main() {
	//resp, err := http.Get("https://www.baidu.com")
	request, err := http.NewRequest(http.MethodGet, "http://www.baidu.com", nil)
	request.Header.Add("User-Agent", "Mozilla/5.0 (Linux; Android 6.0; Nexus 5 Build/MRA58N) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/89.0.4389.90 Mobile Safari/537.36")
	resp, err := http.DefaultClient.Do(request)
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()
	s, err := httputil.DumpResponse(resp, true)
	if err != nil {
		panic(err)
	}
	fmt.Printf("%s", s)
}
