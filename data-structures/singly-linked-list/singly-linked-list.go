package list

import "fmt"

// Node represents an node in the list
type Node struct {
	next  *Node
	Value interface{}
}

// Next returns the next node
func (n *Node) Next() *Node {
	return n.next
}

// SinglyLinkedList represents a Singly Linked List
type SinglyLinkedList struct {
	front  *Node // 头结点
	length int
}

// New returns an initialized list
func New() *SinglyLinkedList {
	return new(SinglyLinkedList).init()
}

func (l *SinglyLinkedList) init() *SinglyLinkedList {
	l.front = nil
	l.length = 0
	return l
}

// Front returns the front node of the list
func (l *SinglyLinkedList) Front() *Node {
	return l.front
}

// Back returns the back node(last node) of the list.
func (l *SinglyLinkedList) Back() *Node {
	if l.front == nil {
		return l.front
	}

	current := l.front
	for current.next != nil {
		current = current.next
	}
	return current
}

// Append add a node to the back of the list
func (l *SinglyLinkedList) Append(n *Node) {
	if l.front == nil {
		l.front = n
		l.length++
		return
	}

	current := l.front
	for current.next != nil {
		current = current.next
	}
	current.next = n
	n.next = nil
	l.length++
}

// Prepend add a node to the front of the list
func (l *SinglyLinkedList) Prepend(n *Node) {
	if l.front == nil {
		l.front = n
	} else {
		n.next = l.front
		l.front = n
	}

	l.length++
}

// InsertBefore inserts the node `n` to before the node `bn`
func (l *SinglyLinkedList) InsertBefore(n, bn *Node) {
	// 1. bn 前面没有结点了，即 bn == front, n.next = bn。 相当于 front !=nil 的 Prepend
	// 2. bn 前面还有结点 bn-1, 即 （bn-1).next = n, n.next = bn
	// 因为需要查找，效率不会比 InsertAfter 高
	if bn == l.front {
		n.next = bn
		l.front = n
	} else {
		current := l.front
		for current.next != nil && current.next != bn {
			current = current.next
		}
		//for now current.next == bn
		current.next = n
		n.next = bn
	}

	l.length++
}

// InsertAfter inserts the node `n` after the node `an`
func (l *SinglyLinkedList) InsertAfter(n, an *Node) {
	// 首先找到 back (front 肯定不等于 nil)

	// an 是 back
	if an.next == nil {
		an.next = n
	} else {
		current := an.next
		an.next = n
		n.next = current
	}
	l.length++
}

// Remove removes a node from the list
func (l *SinglyLinkedList) Remove(n *Node) {
	if l.front == n {
		l.front = n.next
	} else {
		current := l.front
		for current != nil && current.next != n {
			current = current.next
		}
		// for now: current.next == n
		current.next = n.next
	}

	l.length--
}

// GetAtPos return the node in the index `index`
func (l *SinglyLinkedList) GetAtPos(index int) *Node {
	// 从front开始数数
	current := l.front
	count := 0
	for count < index && current != nil && current.next != nil {
		current = current.next
		count++
	}
	if count == index {
		return current
	}
	return nil
}

// Find finds the node with Value equals `value`
func (l *SinglyLinkedList) Find(value interface{}) *Node {
	// 1. front 为空，直接返回 nil
	// 2. front 不为空，value 不存在，返回 nil
	if l.front == nil {
		return nil
	}
	current := l.front

	for current.next != nil {
		fmt.Println("===========>", current.Value, value)
		if current.Value == value {
			return current
		}
		current = current.next
	}

	if current.Value == value {
		return current
	}

	return nil
}

// Length returns length of the list
func (l *SinglyLinkedList) Length() int {
	return l.length
}
