package main

import (
	"bytes"
	"fmt"
	_ "net/http/pprof"
	"runtime"
)

func Greet(writer *bytes.Buffer, name string) {
	fmt.Fprintf(writer, "Hello, %s", name)
	runtime.GC()
}
