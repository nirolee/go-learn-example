package main

import (
	"fmt"
	"gin/queue"
)

func main() {
	q := queue.Queue{1}
	q.Push(2)
	q.Pop()
	q.Pop()
	fmt.Println(q.IsEmpty())
}
