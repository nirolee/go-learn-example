package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

type sp interface {
	Out(key string, val interface{})
	Rd(key string, timeout time.Duration) interface{}
}

type Map struct {
	c   map[string]*entry
	rmx *sync.RWMutex
}

type entry struct {
	ch      chan struct{}
	value   interface{}
	isExist bool
}

func (m *Map) Out(key string, val interface{}) {
	m.rmx.Lock()
	defer m.rmx.Unlock()
	item, ok := m.c[key]
	if !ok {
		m.c[key] = &entry{
			value:   val,
			isExist: true,
		}
		return
	}
	item.value = val
	if !item.isExist {
		if item.ch != nil {
			close(item.ch)
			item.ch = nil
		}
	}
}

func test1() {
	out := make(chan int, 5)
	wg := sync.WaitGroup{}
	wg.Add(2)
	go func() {
		defer wg.Done()
		for i := 0; i < 5; i++ {
			out <- rand.Int()
		}
		close(out)
	}()
	go func() {
		defer wg.Done()
		for i := range out {
			fmt.Println(i)
		}
	}()
	wg.Wait()
}

func test2() {
	ch := make(chan int, 1000)
	go func() {
		for i := 0; i < 10; i++ {
			ch <- i
		}
		close(ch)
	}()
	go func() {
		for {
			a, ok := <-ch
			if !ok {
				fmt.Println("close")
				return
			}
			fmt.Println("a: ", a)
		}
	}()

	fmt.Println("ok")
	time.Sleep(time.Second * 100)
}

type MyMutex struct {
	count int
	sync.Mutex
}

func main() {
	//test1()
	//test2()
}
