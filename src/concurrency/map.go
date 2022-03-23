package concurrency

import "sync"

type Map[K comparable, V any] struct {
	im  map[K]V
	rwm *sync.RWMutex
}

func NewMap[K comparable, V any]() *Map[K, V] {
	m := Map[K, V]{
		im:  map[K]V{},
		rwm: &sync.RWMutex{},
	}

	return &m
}

func (m *Map[K, V]) Write(i map[K]V) {
	m.rwm.Lock()
	for k, v := range i {
		m.im[k] = v
	}
	m.rwm.Unlock()
}

func (m *Map[K, V]) Read(i K) V {
	m.rwm.RLock()
	r := m.im[i]
	m.rwm.RUnlock()
	return r
}
