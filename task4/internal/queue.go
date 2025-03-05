package internal

import (
	"errors"
)

type Queue[T any] struct {
	head *node[T]
	tail *node[T]
}

func NewQueue[T any]() *Queue[T] {
	return &Queue[T]{nil, nil}
}

func (queue *Queue[T]) IsEmpty() bool {
	return queue.head == nil
}

func (queue *Queue[T]) Peek() (T, error) {
	if queue.head == nil {
		return *new(T), errors.New("queue is empty")
	}

	return queue.head.Value, nil
}

func (queue *Queue[T]) Enqueue(value T) {
	newNode := &node[T]{value, nil}

	if queue.tail == nil {
		queue.tail = newNode
		queue.head = newNode
	} else {
		queue.tail.Next = newNode
		queue.tail = newNode
	}
}

func (queue *Queue[T]) Dequeue() (T, error) {
	if queue.head == nil {
		return *new(T), errors.New("queue is empty")
	}

	oldHead := queue.head
	queue.head = queue.head.Next

	if queue.head == nil {
		queue.tail = nil
	}

	return oldHead.Value, nil
}
