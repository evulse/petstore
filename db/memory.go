package db

import (
	"sync"
)

type Record interface {
	GetId() uint64
	SetId(uint64)
}

type MemoryDB struct {
	sync.RWMutex
	records map[string][]Record
}

func NewMemoryDB() *MemoryDB {
	return &MemoryDB{records: map[string][]Record{}}
}

func (m *MemoryDB) CreateTable(table string) {
	m.records[table] = []Record{}
}

func (m *MemoryDB) Insert(table string, record Record) *Record {
	m.Lock()
	defer m.Unlock()
	record.SetId(uint64(len(m.records[table]) + 1))
	m.records[table] = append(m.records[table], record)
	return &record
}

func (m *MemoryDB) Update(table string, record Record) *Record {
	m.Lock()
	defer m.Unlock()
	if record.GetId() > uint64(len(m.records[table]) + 1) {
		return nil
	}
	m.records[table][record.GetId() - 1] = record
	return &record
}

func (m *MemoryDB) Find(table string, id uint64) Record {
	m.RLock()
	defer m.RUnlock()
	if id >= uint64(len(m.records[table]) + 1) {
		return nil
	}
	return m.records[table][id-1]
}

func (m *MemoryDB) FindAll(table string) []Record {
	m.RLock()
	defer m.RUnlock()
	return m.records[table]
}
