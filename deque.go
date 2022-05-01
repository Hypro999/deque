package deque

import (
	"errors"
)

type node[T interface{}] struct {
	val  T
	prev *node[T]
	next *node[T]
}

type Deque[T interface{}] struct {
	head *node[T]
	tail *node[T]
	size uint
}

func (d *Deque[T]) AddFirst(val T) {
	newNode := new(node[T])
	newNode.val = val
	if d.size == 0 {
		d.head = newNode
		d.tail = newNode
	} else {
		newNode.next = d.head
		d.head.prev = newNode
		d.head = newNode
	}
	d.size++
}

func (d *Deque[T]) RemoveFirst() (T, error) {
	if d.size == 0 {
		return *new(T), errors.New("empty deque")
	}
	val := d.head.val
	if d.size == 1 {
		d.head = nil
		d.tail = nil
	} else {
		d.head = d.head.next
		d.head.prev = nil
	}
	d.size--
	return val, nil
}

func (d *Deque[T]) GetFirst() (T, error) {
	if d.size == 0 {
		return *new(T), errors.New("empty deque")
	}
	return d.head.val, nil
}

func (d *Deque[T]) AddLast(val T) {
	newNode := new(node[T])
	newNode.val = val
	if d.size == 0 {
		d.tail = newNode
		d.head = newNode
	} else {
		newNode.prev = d.tail
		d.tail.next = newNode
		d.tail = newNode
	}
	d.size++
}

func (d *Deque[T]) RemoveLast() (T, error) {
	if d.size == 0 {
		return *new(T), errors.New("empty deque")
	}
	val := d.tail.val
	if d.size == 1 {
		d.tail = nil
		d.head = nil
	} else {
		d.tail = d.tail.prev
		d.tail.next = nil
	}
	d.size--
	return val, nil
}

func (d *Deque[T]) GetLast() (T, error) {
	if d.size == 0 {
		return *new(T), errors.New("empty deque")
	}
	return d.tail.val, nil
}
