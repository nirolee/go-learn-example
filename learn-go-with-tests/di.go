package main

import (
	"bytes"
	"fmt"
)

func Greet(write *bytes.Buffer, name string) {
	fmt.Printf("Hello, %s", name)
}
