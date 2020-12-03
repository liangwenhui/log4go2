package core

import (
	"fmt"
	"testing"
)

func TestId(t *testing.T) {
	gen := GetIdGen()
	for i := 0; i < 100; i++ {
		fmt.Print(gen.GetId())
	}
}
