package core

import (
	"bufio"
	"bytes"
	"encoding/json"
	"os"
	"sync"
	"sync/atomic"
)

var C chan []ArgInfo = make(chan []ArgInfo, 10000000)

type ArgInfo struct {
	Seq    int64         `json:"seq"`
	FmtId  int64         `json:"fmt_id"`
	Args   []interface{} `json:"args"`
	ArgLen int           `json:"arg_len"`
}

type ArgBuffer struct {
	sync.RWMutex
	argList []ArgInfo
	maxSize int
	seq     int64
}

func NewDefArgBuffer() *ArgBuffer {
	buffer := &ArgBuffer{
		argList: make([]ArgInfo, 2048),
		maxSize: 2048,
		seq:     0,
	}
	go func() {
		for {
			select {
			case list := <-C:
				buffer.flush(list)
			}
		}
	}()
	return buffer
}

func (buffer *ArgBuffer) Put(info ArgInfo) {
	index := buffer.seq % 2048
	if index == 0 && buffer.seq > 0 {
		a := make([]ArgInfo, 2048)
		copy(a, buffer.argList)
		//buffer.flush(a)
		C <- a
	}
	info.Seq = buffer.seq
	buffer.argList[index] = info
	atomic.AddInt64(&(buffer.seq), 1)
	//buffer.argList = append(buffer.argList,info)
}

func (b *ArgBuffer) flush(a []ArgInfo) {
	buffer := bytes.NewBuffer(make([]byte, 0))
	argFile, _ := os.OpenFile("./info.arg", os.O_CREATE|os.O_APPEND|os.O_RDWR, 0766)
	writer := bufio.NewWriter(argFile)
	for _, v := range a {
		bs, _ := json.Marshal(&v)
		buffer.Write(bs)
		buffer.WriteString("\n")
	}
	writer.Write(buffer.Bytes())

}
