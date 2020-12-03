package core

import (
	"fmt"
	"testing"
)

func TestArgBuffer_Put(t *testing.T) {
	buffer := NewDefArgBuffer()
	buffer.Put(ArgInfo{
		fmtId:  1,
		args:   []interface{}{"", "2", ""},
		argLen: 3,
	})
	fmt.Println(buffer)
}
