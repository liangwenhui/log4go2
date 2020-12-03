package core

import (
	"fmt"
	"testing"
)

func TestLoga_Printf(t *testing.T) {
	loga := NewLoga()
	loga.Printf("aaaa", 1, "2", 3.3)
	loga.Printf("bb", 11, "2b", 34.3)
	loga.Printf("aaaa", 1, "2a", 5.3)
	loga.Printf("cccc", 1, "cc", 6.3)
	fmt.Print("ok")
}
