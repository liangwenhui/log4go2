package core

import (
	"fmt"
	"testing"
)

func Test(t *testing.T) {
	buffer := NewDefFmtBuffer()
	buffer.Put("aaaaaaaa")
	buffer.Put("bbbbb")
	buffer.Put("cccccc")
	buffer.Put("aaaaaaaa")

	fmtA := buffer.GetByFmt("aaaaaaaa")
	fmtB := buffer.GetByFmt("bbbbb")
	fmtC := buffer.GetById(3)

	fmt.Println(fmtA)
	fmt.Println(fmtB)
	fmt.Println(fmtC)

}
