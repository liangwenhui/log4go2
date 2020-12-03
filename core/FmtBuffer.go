package core

import "sync"

const (
	INDEX_FLAG string = "*"
	LEN_FLAG   string = "#"
	FMT_FLAG   string = "$"
)

var idGen *IdGen = GetIdGen()

type FmtCache struct {
	id  int64
	fmt string
	len int
}

type FmtBuffer struct {
	sync.RWMutex
	idCacheMap map[int64]FmtCache
	fmtIdMap   map[string]int64
	maxSize    int
}

func NewDefFmtBuffer() FmtBuffer {
	return FmtBuffer{
		idCacheMap: make(map[int64]FmtCache, 16),
		fmtIdMap:   make(map[string]int64, 16),
	}
}

func (buffer *FmtBuffer) Put(fmt string) {

	buffer.Lock()
	defer buffer.Unlock()
	id := buffer.fmtIdMap[fmt]
	if id == 0 {
		id = idGen.GetId()
		cache := FmtCache{id: id, fmt: fmt, len: len(fmt)}
		buffer.idCacheMap[id] = cache
		buffer.fmtIdMap[fmt] = id
	}

}

func (buffer *FmtBuffer) GetByFmt(fmt string) FmtCache {
	id := buffer.fmtIdMap[fmt]
	return buffer.idCacheMap[id]
}
func (buffer *FmtBuffer) GetById(id int64) FmtCache {
	return buffer.idCacheMap[id]
}
