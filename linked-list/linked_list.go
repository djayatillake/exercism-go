// packaged linkedlist implements a linked list of nodes
package linkedlist

import "errors"

type Node struct {
	Parent *Node
	Val    interface{}
	Child  *Node
}

type List struct {
	first *Node
	last  *Node
}

var ErrEmptyList = errors.New("list is empty")

func (e *Node) Next() *Node {
	if e == nil {
		return nil
	}
	return e.Child
}

func (e *Node) Prev() *Node {
	if e == nil {
		return nil
	}
	return e.Parent
}

func NewList(s ...interface{}) *List {
	l := new(List)
	for _, v := range s {
		l.PushBack(v)
	}
	return l
}

func (l *List) PushFront(v interface{}) {
	n := &Node{Val: v}
	if l.first == nil {
		l.last, l.first = n, n
	} else {
		n.Child = l.first
		l.first.Parent = n
		l.first = n
	}
}

func (l *List) PushBack(v interface{}) {
	n := &Node{Val: v}
	if l.first == nil {
		l.last, l.first = n, n
	} else {
		n.Parent = l.last
		l.last.Child = n
		l.last = n
	}
}

func (l *List) PopBack() (interface{}, error) {
	if l == nil || l.last == nil {
		return nil, ErrEmptyList
	}
	n, val := l.last, l.last.Val
	l.last = n.Parent
	if l.last == nil {
		l.first = nil
	} else {
		l.last.Child = nil
	}
	n = nil
	return val, nil
}

func (l *List) PopFront() (interface{}, error) {
	if l == nil || l.first == nil {
		return nil, ErrEmptyList
	}
	n, val := l.first, l.first.Val
	l.first = n.Child
	if l.first == nil {
		l.first = nil
	} else {
		l.first.Parent = nil
	}
	n = nil
	return val, nil
}

func (l *List) First() *Node {
	return l.first
}

func (l *List) Last() *Node {
	return l.last
}

func (l *List) Reverse() *List {
	for n := l.first; n != nil; n = n.Parent {
		n.Child, n.Parent = n.Parent, n.Child
	}
	l.first, l.last = l.last, l.first
	return l
}
