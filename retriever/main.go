package main

import (
	"fmt"
	"gin/retriever/mock"
	"gin/retriever/real"
	"time"
)

type Retriever interface {
	Get(url string) string
}

type Poster interface {
	Post(url string, form map[string]string) string
}

type RetrieverPoster interface {
	Retriever
	Poster
}

const url = "https://www.baidu.com"

func session(r RetrieverPoster) interface{} {
	r.Post(url, map[string]string{
		"c": "fake",
	})
	return r.Get(url)
}

func download(r Retriever) string {
	return r.Get("http://www.baidu.com")
}

func post(p Poster) string {
	return p.Post("http://www.baidu.com",
		map[string]string{
			"key": "value",
		},
	)
}

func main() {
	var r Retriever
	r = mock.Retriever{"fake"}
	r = &real.Retriever{
		UserAgent: "Mozilla/5.0",
		TimeOut:   time.Minute,
	}
	inspect(r)
	if mockRetriever, ok := r.(*mock.Retriever); ok {
		fmt.Println(mockRetriever.Contents)
	} else {
		fmt.Println("not a mock retriever")
	}

}

func inspect(r Retriever) {
	fmt.Println("Inspecting", r)
	fmt.Printf("> %T %v\n", r, r)
	fmt.Print(" > Type switch：")
	switch v := r.(type) {
	case mock.Retriever:
		fmt.Println("Content:", v.Contents)
	case *real.Retriever:
		fmt.Println("UserAgent：", v.UserAgent)
	}

}
