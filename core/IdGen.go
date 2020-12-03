package core

import (
	"reflect"
	"sync"
	"sync/atomic"
)

var idgen IdGen
var lock sync.Mutex = sync.Mutex{}

type IdGen struct {
	index int64
}

/*
获取id
*/
func (gen *IdGen) GetId() int64 {
	atomic.AddInt64(&gen.index, 1)
	return gen.index
}

func (x IdGen) IsStructureEmpty() bool {
	return reflect.DeepEqual(x, IdGen{})
}

func GetIdGen() *IdGen {
	if idgen.IsStructureEmpty() {
		lock.Lock()
		if idgen.IsStructureEmpty() {
			idgen = IdGen{0}
		}
		lock.Unlock()
	}
	return &idgen
}
