package core

import (
	"bufio"
	"bytes"
	"os"
	"strconv"
)

type Loga struct {
	fmtBuffer *FmtBuffer
	argBuffer *ArgBuffer
}

func NewLoga() Loga {
	return Loga{
		fmtBuffer: NewDefFmtBuffer(),
		argBuffer: NewDefArgBuffer(),
	}
}

func (l *Loga) Printf(fmt string, args ...interface{}) {
	fmtId := l.fmtBuffer.Put(fmt)
	l.argBuffer.Put(ArgInfo{
		Args:   args,
		ArgLen: len(args),
		FmtId:  fmtId,
	})
}

func (l *Loga) Flush() {
	fb := l.fmtBuffer
	fmtFile, _ := os.OpenFile("./info.fmt", os.O_CREATE|os.O_RDWR, 0766)
	buffer := bytes.NewBuffer(make([]byte, 0))
	writer := bufio.NewWriter(fmtFile)
	for k, v := range fb.idCacheMap {
		buffer.WriteString(INDEX_FLAG + strconv.FormatInt(k, 10) + LEN_FLAG + strconv.Itoa(v.len) + FMT_FLAG + v.fmt)
	}
	writer.Write(buffer.Bytes())
	buffer.Reset()
	ab := l.argBuffer
	index := ab.seq % 2048
	//argFile, _ := os.OpenFile("./info.arg",os.O_CREATE|os.O_APPEND|os.O_RDWR, 0766)
	//writer = bufio.NewWriter(argFile)
	a := make([]ArgInfo, index)
	copy(a, ab.argList)
	C <- a

}
