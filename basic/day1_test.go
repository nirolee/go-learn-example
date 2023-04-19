package basic

import (
	"bufio"
	"bytes"
	"testing"
)

func TestFirst(t *testing.T) {
	first()
}

func TestSecond(t *testing.T) {
	input := "niro"
	reader := bufio.NewReader(bytes.NewBufferString(input))
	second(reader)
	input = "29"
	reader = bufio.NewReader(bytes.NewBufferString(input))
	second(reader)
}

func TestThird(t *testing.T) {
	third(10)
}
