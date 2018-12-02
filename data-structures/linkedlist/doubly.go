package linkedlist

import (
	"errors"
	"reflect"
)

// DoublyLinkedList ...
type DoublyLinkedList struct {
	head *Node
	tail *Node
}

// PushFront add to front
func (s *DoublyLinkedList) PushFront(key Key) {
	s.head = &Node{
		Key:  key,
		Next: s.head,
	}

	if s.tail == nil {
		s.tail = s.head
	}
}

// TopFront return front item
func (s *DoublyLinkedList) TopFront() Key {
	return s.head.Key
}

// PopFront remove front item
func (s *DoublyLinkedList) PopFront() error {

	if s.head == nil {
		return errors.New("empty list")
	}

	if s.head = s.head.Next; s.head == nil {
		s.tail = nil
	}

	return nil
}

// PushBack add to back
func (s *DoublyLinkedList) PushBack(key Key) {
	node := &Node{
		Key: key,
	}

	if s.tail == nil {
		s.tail = node
		s.head = s.tail
	} else {
		s.tail.Next = node
		s.tail.Prev = s.tail
		s.tail = node
	}
}

// TopBack return back item
func (s *DoublyLinkedList) TopBack() Key {
	return s.tail.Key
}

// PopBack remove back item
func (s *DoublyLinkedList) PopBack() error {

	if s.head == nil {
		return errors.New("empty list")
	}

	if reflect.DeepEqual(s.head, s.tail) {
		s.tail = nil
		s.head = s.tail

	} else {
		s.tail = s.tail.Prev
		s.tail.Next = nil
	}

	return nil
}

// Empty empty list?
func (s *DoublyLinkedList) Empty() bool {
	if s.head != nil {
		return false
	}

	return true
}

// Find is key in list?
func (s *DoublyLinkedList) Find(key Key) bool {

	if s.head == nil {
		return false
	}

	if s.head.Key == key || s.tail.Key == key {
		return true
	}

	iterator := s.head

	for iterator.Next.Next != nil {
		if iterator = iterator.Next; iterator.Key == key {
			return true
		}
	}

	return false
}

// Erase remove key from list
func (s *DoublyLinkedList) Erase(key Key) error {
	if s.head == nil {
		return errors.New("empty list")
	}

	if s.head.Key == key {
		s.head = nil
	} else {
		iterator := s.head

		for iterator.Next != nil {
			if iterator.Key != key {
				iterator = iterator.Next
			} else {
				iterator.Next = nil
			}
		}

		s.head = iterator.Next
	}

	return nil
}
