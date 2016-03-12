package stack

import (
	"testing"
)

func TestStack(t *testing.T) {
	s := NewArrayImplements()
	if !s.IsEmpty() {
		t.Errorf("stack should be empty")
	}
	s.Push(1)
	s.Push(2)
	s.Push(3)

	if s.IsEmpty() {
		t.Errorf("should not be empty")
	}

	if s.Peek() != 3 {
		t.Errorf("got %v, should be %v", s.Peek(), 3)
	}

	if s.Pop() != 3 {
		t.Errorf("go %v, should be %v", s.Peek(), 3)
	}
	if s.Peek() != 2 {
		t.Errorf("go %v, should be %v", s.Peek(), 2)
	}
	if s.Len() != 2 {
		t.Errorf("go %v, should be %v", s.Len(), 2)
	}
}

func BenchmarkPush(b *testing.B) {
	s := NewArrayImplements()
	for i := 0; i < b.N; i++ {
		s.Push(i)
	}
}

func BenchmarkPop(b *testing.B) {
	s := NewArrayImplements()
	for i := 0; i < b.N; i++ {
		s.Push(i)
	}

	b.ResetTimer()
	b.Log(s.Peek())
	for !s.IsEmpty() {
		s.Pop()
	}
}
