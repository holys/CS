package stack

import "sync"

type Stack interface {
	Push(item interface{})
	Pop() (item interface{})
	IsEmpty() bool
	Peak() interface{}
	Len() int // unlimited?
}

var _ Stack = (*arrayImplements)(nil)

type arrayImplements struct {
	items []interface{}
	lock  sync.Mutex
	len   int
}

func NewArrayImplements() Stack {
	s := new(arrayImplements)
	s.items = make([]interface{}, 0)
	s.len = 0
	return s
}

func (s *arrayImplements) Push(item interface{}) {
	s.lock.Lock()
	s.items = append(s.items, item)
	s.len++
	s.lock.Unlock()
}

func (s *arrayImplements) Pop() (item interface{}) {
	s.lock.Lock()
	defer s.lock.Unlock()
	item, s.items = s.items[s.len-1], s.items[:s.len-1]
	s.len--
	return item
}

func (s *arrayImplements) IsEmpty() bool {
	s.lock.Lock()
	defer s.lock.Unlock()
	return s.len == 0
}

func (s *arrayImplements) Peak() interface{} {
	s.lock.Lock()
	defer s.lock.Unlock()
	return s.items[s.len-1]
}

func (s *arrayImplements) Len() int {
	s.lock.Lock()
	defer s.lock.Unlock()
	return s.len
}
