package linkedlist

import (
	"errors"
)

type Node struct {
	Value string
	Next  *Node
	Prev  *Node
}

func newNode(v string) Node {
	return Node{Value: v, Next: nil, Prev: nil}
}

type LinkedList struct {
	Start *Node
	Tail  *Node
	Len   uint
}

func NewLinkedList() LinkedList {
	return LinkedList{Start: nil, Tail: nil, Len: 0}
}

func (ll *LinkedList) AddNode(v string) {
	newNode := newNode(v)

	if ll.Len == 0 {
		ll.Start = &newNode
	} else if ll.Len == 1 {
		ll.Start.Next = &newNode
		newNode.Prev = ll.Start
		ll.Tail = &newNode
	} else {
		ll.Tail.Next = &newNode
		newNode.Prev = ll.Tail
		ll.Tail = &newNode
	}

	ll.Len += 1
}

func (ll *LinkedList) Get(i uint) (string, error) {
	if i >= ll.Len {
		return "", errors.New("out of bounds index")
	}

	if i == 0 {
		return ll.Start.Value, nil
	} else if i == ll.Len-1 {
		return ll.Tail.Value, nil
	}

	pivot := ll.Start
	for ; i != 0; i-- {
		pivot = pivot.Next
	}

	return pivot.Value, nil
}

func (ll *LinkedList) Set(i uint, v string) error {
	if i >= ll.Len {
		return errors.New("out of bounds index")
	}

	if i == 0 {
		ll.Start.Value = v
	} else if i == ll.Len-1 {
		ll.Tail.Value = v
	}

	pivot := ll.Start
	for ; i != 0; i-- {
		pivot = pivot.Next
	}
	pivot.Value = v

	return nil
}

func (ll *LinkedList) Remove(i uint) error {
	if i >= ll.Len {
		return errors.New("out of bounds index")
	}

	// Remover o primeiro elemento
	if i == 0 {
		// Remover o primeiro elemento com lista com 1 elemento
		if ll.Len == 1 {
			ll.Start = nil
		} else { // Remover o primeiro com lista com mais de 1 elemento
			ll.Start = ll.Start.Next
		}
	} else if i == ll.Len-1 { // Remover o ultimo elemento
		sup := ll.Tail
		ll.Tail.Prev.Next = nil
		ll.Tail = sup.Prev
	} else { // Remove elemento no meio
		pivot := ll.Start

		for ; i != 0; i-- {
			pivot = pivot.Next
		}

		pivot.Prev.Next = pivot.Next
		pivot.Next.Prev = pivot.Prev
	}

	ll.Len -= 1
	return nil
}

func (ll *LinkedList) Insert(i uint, v string) error {
	if i >= ll.Len {
		return errors.New("out of bounds index")
	}

	if ll.Len == 0 { // Primeiro com lista vazia
		ll.AddNode(v)
	} else if i == 0 { // Primeiro com lista povoada
		n := newNode(v)
		n.Next = ll.Start
		ll.Start = &n
	} else { // Adicionar no meio
		n := newNode(v)

		pivot := ll.Start
		for i != 1 {
			pivot = pivot.Next
		}

		n.Next = pivot.Next
		n.Prev = pivot
		pivot.Next.Prev = &n
		pivot.Next = &n
	}

	ll.Len += 1
	return nil
}
