package db

import "github.com/huandu/skiplist"

type MemTable struct {
	table skiplist.SkipList
}

func (mt *MemTable) Set(key []byte, value []byte) {
	mt.table.Set(key, value)
}

func (mt *MemTable) Delete(key []byte) {
	mt.table.Remove(key)
}

func (mt *MemTable) Get(key []byte) ([]byte, bool) {
	elem := mt.table.Get(key)
	if elem != nil {
		return elem.Value.([]byte), true
	} else {
		return nil, false
	}
}
