package libs

import "sync"

// globalImports stores []ImportEntry in a place where GC won't
// delete it
type SafeMap struct {
	sync.RWMutex
	idx int
	m   map[int]interface{}
}

func (s *SafeMap) nextidx() int {
	s.Lock()
	defer s.Unlock()
	s.idx++
	return s.idx
}

func (s *SafeMap) init() {
	s.m = make(map[int]interface{})
}

func (s *SafeMap) Get(idx int) interface{} {
	s.RLock()
	defer s.RUnlock()
	return s.m[idx]
}

func (s *SafeMap) Del(idx int) {
	s.Lock()
	delete(s.m, idx)
	s.Unlock()
}

// set accepts an entry and returns an index for it
func (s *SafeMap) Set(ie interface{}) int {
	idx := s.nextidx()
	s.Lock()
	s.m[idx] = ie
	defer s.Unlock()

	return idx
}
