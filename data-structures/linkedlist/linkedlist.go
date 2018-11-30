package linkedlist

// Key ...
type Key interface{}

// LinkedList ...
type LinkedList interface {
	PushFront(Key)
	TopFront() Key
	PopFront() error
	PushBack(Key)
	TopBack() Key
	PopBack() error
	Empty() bool
	Erase(Key) error
	Find(Key) bool
}

// Node ...
type Node struct {
	Key
	Next *Node
	Prev *Node
}
