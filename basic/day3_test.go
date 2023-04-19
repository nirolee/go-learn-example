package basic

import (
	"testing"
)

func TestDogRun(t *testing.T) {
	dog := Dog{name: "dog"}
	dog.run()
}
