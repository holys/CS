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
	l := New()
	if l.Front() != nil {
		t.Error("Expected nil, got ", l.Front())
	}

	n1 := &Node{Value: "test1"}
	l.Append(n1)
	if l.Front() != n1 {
		t.Error("Expected ", n1, ", got ", l.Front())
	}

	n2 := &Node{Value: "test2"}
	l.Append(n2)
	if l.Front() != n1 {
		t.Error("Expected ", n1, ", got ", l.Front())
	}
}

func TestBack(t *testing.T) {
	l := New()
	if l.Back() != nil {
		t.Error("Expected nil, got ", l.Back())
	}

	n1 := &Node{Value: "test1"}
	l.Append(n1)
	if l.Back() != n1 {
		t.Error("Expected ", n1, ", got ", l.Back())
	}

	n2 := &Node{Value: "test2"}
	l.Append(n2)
	if l.Back() != n2 {
		t.Error("Expected ", n2, ", got ", l.Back())
	}
}

func TestAppend(t *testing.T) {
	l := New()
	n1 := &Node{Value: "test1"}
	n2 := &Node{Value: "test2"}
	n3 := &Node{Value: "test3"}

	if l.Length() != 0 {
		t.Error("Expected 0, got ", l.Length())
	}
	if l.front != nil {
		t.Error("Expected nil, got ", l.front)
	}

	l.Append(n1)
	if l.Length() != 1 {
		t.Error("Expected 1, got ", l.Length())
	}
	if l.front != n1 {
		t.Error("Expected ", n1, ", got ", l.front)
	}

	l.Append(n2)
	if l.Length() != 2 {
		t.Error("Expected 2, got ", l.Length())
	}
	if l.front.next != n2 {
		t.Error("Expected ", n2, ", got ", l.front.next)
	}

	l.Append(n3)
	if l.Length() != 3 {
		t.Error("Expected 3, got ", l.Length())
	}
	if l.front.next.next != n3 {
		t.Error("Expected ", n3, ", got ", l.front.next.next)
	}

	l.init()
	for i := 0; i < 5; i++ {
		l.Append(&Node{Value: i})
		if l.front.Value != 0 {
			t.Errorf("expected 0 got %v", l.front.Value)
		}
	}
}

func TestPrepend(t *testing.T) {
	l := New()
	n1 := &Node{Value: "test1"}
	n2 := &Node{Value: "test2"}

	l.Prepend(n1)
	if l.Length() != 1 {
		t.Error("Expected 1, got ", l.Length())
	}
	if l.front != n1 {
		t.Error("Expected ", n1, ", got ", l.front)
	}

	l.Prepend(n2)
	if l.Length() != 2 {
		t.Error("Expected 2, got ", l.Length())
	}
	if l.front != n2 {
		t.Error("Expected ", n2, ", got ", l.front)
	}
	if l.front.next != n1 {
		t.Error("Expected ", n1, ", got ", l.front.next)
	}

	for i := 0; i < 5; i++ {
		n := &Node{Value: i}
		l.Prepend(n)
		if l.front != n {
			t.Errorf("expected %v got %v", n, l.front)
		}
	}
}

func TestInsertBefore(t *testing.T) {
	l := New()
	n1 := &Node{Value: "test1"}
	n2 := &Node{Value: "test2"}

	l.Append(n1)
	l.InsertBefore(n2, n1)
	if l.Length() != 2 {
		t.Error("Expected 2, got ", l.Length())
	}
	if l.front != n2 {
		t.Error("Expected ", n2, ", got ", l.front)
	}
	if l.front.next != n1 {
		t.Error("Expected ", n1, ", got ", l.front.next)
	}

	n3 := &Node{Value: "test3"}
	l.InsertBefore(n3, n1)
	if l.Length() != 3 {
		t.Error("Expected 3, got ", l.Length())
	}
	if l.front.next != n3 {
		t.Error("Expected ", n3, ", got ", l.front.next)
	}
	if l.front.next.next != n1 {
		t.Error("Expected ", n1, ", got ", l.front.next.next)
	}

	n4 := &Node{Value: "test4"}
	l.InsertBefore(n4, n1)
	if l.Length() != 4 {
		t.Error("Expected 4, got ", l.Length())
	}
	if l.front.next.next != n4 {
		t.Error("Expected ", n4, ", got ", l.front.next.next)
	}
	if l.front.next.next.next != n1 {
		t.Error("Expected ", n1, ", got ", l.front.next.next.next)
	}
}

func TestInsertAfter(t *testing.T) {
	l := New()
	n1 := &Node{Value: "test1"}
	n2 := &Node{Value: "test2"}

	l.Append(n1)
	l.InsertAfter(n2, n1)
	v := l.front.next
	if l.Length() != 2 {
		t.Error("Expected 2, got ", l.Length())
	}
	if v != n2 {
		t.Error("Expected ", n2, ", got ", v)
	}

	n3 := &Node{Value: "test3"}
	l.InsertAfter(n3, n2)
	v = l.front.next.next
	if l.Length() != 3 {
		t.Error("Expected 3, got ", l.Length())
	}
	if v != n3 {
		t.Error("Expected ", n3, ", got ", v)
	}
}

// TestRemoveOne tests removing the only node from a list
func TestRemoveOne(t *testing.T) {
	l := New()
	n1 := &Node{Value: "test1"}

	l.Append(n1)
	l.Remove(n1)

	if l.Length() != 0 {
		t.Error("Expected 0, got ", l.Length())
	}
	if l.front != nil {
		t.Error("Expected nil, got ", l.front)
	}
}

// TestRemoveTwofront tests removing the front node of a list with 2 nodes
func TestRemoveTwofront(t *testing.T) {
	l := New()
	n1 := &Node{Value: "test1"}
	n2 := &Node{Value: "test2"}

	l.Append(n1)
	l.Append(n2)
	l.Remove(n1)

	if l.Length() != 1 {
		t.Error("Expected 1, got ", l.Length())
	}
	if l.front != n2 {
		t.Error("Expected ", n2, ", got ", l.front)
	}
}

// TestRemoveTwoLast tests removing the last node of a list with 2 nodes
func TestRemoveTwoLast(t *testing.T) {
	l := New()
	n1 := &Node{Value: "test1"}
	n2 := &Node{Value: "test2"}

	l.Append(n1)
	l.Append(n2)
	l.Remove(n2)

	if l.front.next != nil {
		t.Error("Expected nil, got ", l.front.next)
	}
}

// TestRemoveThreeLast tests removing the last node of a list with 3 nodes
func TestRemoveThreeLast(t *testing.T) {
	l := New()
	n1 := &Node{Value: "test1"}
	n2 := &Node{Value: "test2"}
	n3 := &Node{Value: "test3"}

	l.Append(n1)
	l.Append(n2)
	l.Append(n3)
	l.Remove(n3)

	if l.Length() != 2 {
		t.Error("Expected 2, got ", l.Length())
	}
	if l.front.next.next != nil {
		t.Error("Expected nil, got ", l.front.next.next)
	}
}

// TestRemoveNonExistingNode tests removing a node not found in the list
func TestRemoveNonExistingNode(t *testing.T) {
	l := New()
	n1 := &Node{Value: "test1"}
	n2 := &Node{Value: "test2"}

	// when list is nil
	l.Remove(n1)

	l.Append(n1)
	l.Remove(n2)

	if l.Length() != 1 {
		t.Error("Expected 1, got ", l.Length())
	}
	if l.front != n1 {
		t.Error("Expected ", n1, ", got ", l.front)
	}
}

func TestGetAtPos(t *testing.T) {
	l := New()

	v := l.GetAtPos(3)
	if v != nil {
		t.Error("Expected nil, got ", v)
	}

	n1 := &Node{Value: "test1"}
	l.Append(n1)
	v = l.GetAtPos(0)
	if v != n1 {
		t.Error("Expected ", n1, ", got ", v)
	}

	n2 := &Node{Value: "test2"}
	l.Append(n2)
	v = l.GetAtPos(1)
	if v != n2 {
		t.Error("Expected ", n2, ", got ", v)
	}

	v = l.GetAtPos(2)
	if v != nil {
		t.Errorf("Expected nil, got %v", v)
	}
}

func TestFind(t *testing.T) {
	l := New()

	// list is nil
	v := l.Find("test")
	if v != nil {
		t.Error("Expected nil, got ", v)
	}

	n1 := &Node{Value: "test1"}
	l.Append(n1)

	v = l.Find("test1")
	if v != n1 {
		t.Error("Expected ", n1, ", got ", v)
	}

	n2 := &Node{Value: "test2"}
	l.Append(n2)
	v = l.Find("test2")
	if v != n2 {
		t.Error("Expected ", n2, ", got ", v)
	}
}

func TestFindNoneExistingNode(t *testing.T) {
	l := New()

	for i := 0; i < 10; i++ {
		l.Append(&Node{Value: i})
	}

	found := l.Find(11)
	if found != nil {
		t.Errorf("should be nil, but got %v", found)
	}
}
