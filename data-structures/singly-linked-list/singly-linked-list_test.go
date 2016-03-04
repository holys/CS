package list

import "testing"

// func TestFront(t *testing.T) {

// }

// func TestBack(t *testing.T) {

// }

// func TestAppend(t *testing.T) {
// 	l := New()
// 	n1 := &Node{Value: 1}
// 	n2 := &Node{Value: 2}
// 	l.Append(n1)
// 	l.Append(n2)

// 	if l.length != 2 {
// 		t.Error("should be %d, got %d", 2, l.length)
// 	}

// 	l.init()

// 	for i := 0; i < 10000; i++ {
// 		l.Append(&Node{Value: i})
// 	}
// 	if l.length != 10000 {
// 		t.Errorf("should be 10000, got %d", l.length)
// 	}

// 	l.init()
// 	for i := 0; i < 10000; i++ {
// 		l.Prepend(&Node{Value: i})
// 	}
// 	if l.length != 10000 {
// 		t.Errorf("should be 10000, got %d", l.length)
// 	}
// }

// func TestPrepend(t *testing.T) {

// }

// func TestInsertBefore(t *testing.T) {

// }

// func TestInsertAfter(t *testing.T) {

// }

func BenchmarkAppend(b *testing.B) {
	l := New()

	for i := 0; i < b.N; i++ {
		l.Append(&Node{Value: i})
	}
}

// 比 append 快50倍？ 因为不需要查找
func BenchmarkPrepend(b *testing.B) {
	l := New()

	for i := 0; i < b.N; i++ {
		l.Prepend(&Node{Value: i})
	}
}

func TestFront(t *testing.T) {
	s := New()
	if s.Front() != nil {
		t.Error("Expected nil, got ", s.Front())
	}

	n1 := &Node{Value: "test1"}
	s.Append(n1)
	if s.Front() != n1 {
		t.Error("Expected ", n1, ", got ", s.Front())
	}

	n2 := &Node{Value: "test2"}
	s.Append(n2)
	if s.Front() != n1 {
		t.Error("Expected ", n1, ", got ", s.Front())
	}
}

func TestBack(t *testing.T) {
	s := New()
	if s.Back() != nil {
		t.Error("Expected nil, got ", s.Back())
	}

	n1 := &Node{Value: "test1"}
	s.Append(n1)
	if s.Back() != n1 {
		t.Error("Expected ", n1, ", got ", s.Back())
	}

	n2 := &Node{Value: "test2"}
	s.Append(n2)
	if s.Back() != n2 {
		t.Error("Expected ", n2, ", got ", s.Back())
	}
}

func TestAppend(t *testing.T) {
	s := New()
	n1 := &Node{Value: "test1"}
	n2 := &Node{Value: "test2"}
	n3 := &Node{Value: "test3"}

	if s.Length() != 0 {
		t.Error("Expected 0, got ", s.Length())
	}
	if s.front != nil {
		t.Error("Expected nil, got ", s.front)
	}

	s.Append(n1)
	if s.Length() != 1 {
		t.Error("Expected 1, got ", s.Length())
	}
	if s.front != n1 {
		t.Error("Expected ", n1, ", got ", s.front)
	}

	s.Append(n2)
	if s.Length() != 2 {
		t.Error("Expected 2, got ", s.Length())
	}
	if s.front.next != n2 {
		t.Error("Expected ", n2, ", got ", s.front.next)
	}

	s.Append(n3)
	if s.Length() != 3 {
		t.Error("Expected 3, got ", s.Length())
	}
	if s.front.next.next != n3 {
		t.Error("Expected ", n3, ", got ", s.front.next.next)
	}
}

func TestPrepend(t *testing.T) {
	s := New()
	n1 := &Node{Value: "test1"}
	n2 := &Node{Value: "test2"}

	s.Prepend(n1)
	if s.Length() != 1 {
		t.Error("Expected 1, got ", s.Length())
	}
	if s.front != n1 {
		t.Error("Expected ", n1, ", got ", s.front)
	}

	s.Prepend(n2)
	if s.Length() != 2 {
		t.Error("Expected 2, got ", s.Length())
	}
	if s.front != n2 {
		t.Error("Expected ", n2, ", got ", s.front)
	}
	if s.front.next != n1 {
		t.Error("Expected ", n1, ", got ", s.front.next)
	}
}

func TestInsertBefore(t *testing.T) {
	s := New()
	n1 := &Node{Value: "test1"}
	n2 := &Node{Value: "test2"}

	s.Append(n1)
	s.InsertBefore(n2, n1)
	if s.Length() != 2 {
		t.Error("Expected 2, got ", s.Length())
	}
	if s.front != n2 {
		t.Error("Expected ", n2, ", got ", s.front)
	}
	if s.front.next != n1 {
		t.Error("Expected ", n1, ", got ", s.front.next)
	}

	n3 := &Node{Value: "test3"}
	s.InsertBefore(n3, n1)
	if s.Length() != 3 {
		t.Error("Expected 3, got ", s.Length())
	}
	if s.front.next != n3 {
		t.Error("Expected ", n3, ", got ", s.front.next)
	}
	if s.front.next.next != n1 {
		t.Error("Expected ", n1, ", got ", s.front.next.next)
	}

	n4 := &Node{Value: "test4"}
	s.InsertBefore(n4, n1)
	if s.Length() != 4 {
		t.Error("Expected 4, got ", s.Length())
	}
	if s.front.next.next != n4 {
		t.Error("Expected ", n4, ", got ", s.front.next.next)
	}
	if s.front.next.next.next != n1 {
		t.Error("Expected ", n1, ", got ", s.front.next.next.next)
	}
}

func TestInsertAfter(t *testing.T) {
	s := New()
	n1 := &Node{Value: "test1"}
	n2 := &Node{Value: "test2"}

	s.Append(n1)
	s.InsertAfter(n2, n1)
	v := s.front.next
	if s.Length() != 2 {
		t.Error("Expected 2, got ", s.Length())
	}
	if v != n2 {
		t.Error("Expected ", n2, ", got ", v)
	}

	n3 := &Node{Value: "test3"}
	s.InsertAfter(n3, n2)
	v = s.front.next.next
	if s.Length() != 3 {
		t.Error("Expected 3, got ", s.Length())
	}
	if v != n3 {
		t.Error("Expected ", n3, ", got ", v)
	}
}

// TestRemoveOne tests removing the only node from a list
func TestRemoveOne(t *testing.T) {
	s := New()
	n1 := &Node{Value: "test1"}

	s.Append(n1)
	s.Remove(n1)

	if s.Length() != 0 {
		t.Error("Expected 0, got ", s.Length())
	}
	if s.front != nil {
		t.Error("Expected nil, got ", s.front)
	}
}

// TestRemoveTwofront tests removing the front node of a list with 2 nodes
func TestRemoveTwofront(t *testing.T) {
	s := New()
	n1 := &Node{Value: "test1"}
	n2 := &Node{Value: "test2"}

	s.Append(n1)
	s.Append(n2)
	s.Remove(n1)

	if s.Length() != 1 {
		t.Error("Expected 1, got ", s.Length())
	}
	if s.front != n2 {
		t.Error("Expected ", n2, ", got ", s.front)
	}
}

// TestRemoveTwoLast tests removing the last node of a list with 2 nodes
func TestRemoveTwoLast(t *testing.T) {
	s := New()
	n1 := &Node{Value: "test1"}
	n2 := &Node{Value: "test2"}

	s.Append(n1)
	s.Append(n2)
	s.Remove(n2)

	if s.front.next != nil {
		t.Error("Expected nil, got ", s.front.next)
	}
}

// TestRemoveThreeLast tests removing the last node of a list with 3 nodes
func TestRemoveThreeLast(t *testing.T) {
	s := New()
	n1 := &Node{Value: "test1"}
	n2 := &Node{Value: "test2"}
	n3 := &Node{Value: "test3"}

	s.Append(n1)
	s.Append(n2)
	s.Append(n3)
	s.Remove(n3)

	if s.Length() != 2 {
		t.Error("Expected 2, got ", s.Length())
	}
	if s.front.next.next != nil {
		t.Error("Expected nil, got ", s.front.next.next)
	}
}

// TestRemoveNonExistingNode tests removing a node not found in the list
func TestRemoveNonExistingNode(t *testing.T) {
	s := New()
	n1 := &Node{Value: "test1"}
	n2 := &Node{Value: "test2"}

	s.Append(n1)
	s.Remove(n2)

	if s.Length() != 1 {
		t.Error("Expected 1, got ", s.Length())
	}
	if s.front != n1 {
		t.Error("Expected ", n1, ", got ", s.front)
	}
}

// func TestRemoveBeforeEmpty(t *testing.T) {
// 	s := New()
// 	n1 := &Node{Value: "test1"}

// 	s.RemoveBefore(n1)
// 	v := s.front
// 	if v != nil {
// 		t.Error("Expected nil, got ", v)
// 	}
// }

// func TestRemoveBeforeOneNode(t *testing.T) {
// 	s := New()
// 	n1 := &Node{Value: "test1"}

// 	s.Append(n1)
// 	s.RemoveBefore(n1)

// 	v := s.front
// 	if v != n1 {
// 		t.Error("Expected ", n1, ", got ", v)
// 	}
// }

// func TestRemoveBeforeTwoNodes(t *testing.T) {
// 	s := New()
// 	n1 := &Node{Value: "test1"}
// 	n2 := &Node{Value: "test2"}

// 	s.Append(n1)
// 	s.Append(n2)
// 	s.RemoveBefore(n2)

// 	v := s.front
// 	if v != n2 {
// 		t.Error("Expected ", n2, ", got ", v)
// 	}
// }

// func TestRemoveBeforeThreeNodes(t *testing.T) {
// 	s := New()
// 	n1 := &Node{Value: "test1"}
// 	n2 := &Node{Value: "test2"}
// 	n3 := &Node{Value: "test3"}

// 	s.Append(n1)
// 	s.Append(n2)
// 	s.Append(n3)
// 	s.RemoveBefore(n3)

// 	v := s.front.next
// 	if v != n3 {
// 		t.Error("Expected ", n3, ", got ", v)
// 	}
// }

// func TestRemoveBeforeFourNodes(t *testing.T) {
// 	s := New()
// 	n1 := &Node{Value: "test1"}
// 	n2 := &Node{Value: "test2"}
// 	n3 := &Node{Value: "test3"}
// 	n4 := &Node{Value: "test4"}

// 	s.Append(n1)
// 	s.Append(n2)
// 	s.Append(n3)
// 	s.Append(n4)
// 	s.RemoveBefore(n4)

// 	v := s.front.next.next
// 	if v != n4 {
// 		t.Error("Expected ", n4, ", got ", v)
// 	}
// }

// func TestRemoveAfterEmpty(t *testing.T) {
// 	s := New()
// 	n1 := &Node{Value: "test1"}

// 	s.RemoveAfter(n1)

// 	v := s.front
// 	if v != nil {
// 		t.Error("Expected nil, got ", v)
// 	}
// }

// func TestRemoveAfterOneNode(t *testing.T) {
// 	s := New()
// 	n1 := &Node{Value: "test1"}

// 	s.Append(n1)
// 	s.RemoveAfter(n1)

// 	v := s.front.next
// 	if v != nil {
// 		t.Error("Expected nil, got ", v)
// 	}
// }

// func TestRemoveAfterTwoNodes(t *testing.T) {
// 	s := New()
// 	n1 := &Node{Value: "test1"}
// 	n2 := &Node{Value: "test2"}

// 	s.Append(n1)
// 	s.Append(n2)
// 	s.RemoveAfter(n1)

// 	v := s.front.next
// 	if v != nil {
// 		t.Error("Expected nil, got ", v)
// 	}
// }

// func TestRemoveAfterThreeNodes(t *testing.T) {
// 	s := New()
// 	n1 := &Node{Value: "test1"}
// 	n2 := &Node{Value: "test2"}
// 	n3 := &Node{Value: "test3"}

// 	s.Append(n1)
// 	s.Append(n2)
// 	s.Append(n3)
// 	s.RemoveAfter(n2)

// 	v := s.front.next.next
// 	if v != nil {
// 		t.Error("Expected nil, got ", v)
// 	}
// }

func TestGetAtPos(t *testing.T) {
	s := New()

	v := s.GetAtPos(3)
	if v != nil {
		t.Error("Expected nil, got ", v)
	}

	n1 := &Node{Value: "test1"}
	s.Append(n1)
	v = s.GetAtPos(0)
	if v != n1 {
		t.Error("Expected ", n1, ", got ", v)
	}

	n2 := &Node{Value: "test2"}
	s.Append(n2)
	v = s.GetAtPos(1)
	if v != n2 {
		t.Error("Expected ", n2, ", got ", v)
	}

	v = s.GetAtPos(2)
	if v != nil {
		t.Errorf("Expected nil, got %v", v)
	}
}

func TestFind(t *testing.T) {
	s := New()

	v := s.Find("test")
	if v != nil {
		t.Error("Expected nil, got ", v)
	}

	n1 := &Node{Value: "test1"}
	s.Append(n1)
	v = s.Find("test1")
	if v != n1 {
		t.Error("Expected ", n1, ", got ", v)
	}

	// n2 := &Node{Value: "test2"}
	// s.Append(n2)
	// v = s.Find("test2")
	// if v != n2 {
	// 	t.Error("Expected ", n2, ", got ", v)
	// }
}

func TestFindNoneExistingNode(t *testing.T) {
	l := New()
	// n1 := &Node{Value: 1}

	for i := 0; i < 10; i++ {
		l.Append(&Node{Value: i})
	}

	found := l.Find(11)
	if found != nil {
		t.Errorf("should be nil, but got %v", found)
	}
}
