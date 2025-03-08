package internal

import (
	"errors"
)

type Stack[T any] struct {
	head *node[T]
}

func NewStack[T any]() *Stack[T] {
	return &Stack[T]{nil}
}

func (stack *Stack[T]) IsEmpty() bool {
	return stack.head == nil
}

func (stack *Stack[T]) Peek() (T, error) {
	if stack.head == nil {
		return *new(T), errors.New("stack is empty")
	}

	return stack.head.value, nil
}

func (stack *Stack[T]) Push(value T) {
	stack.head = &node[T]{value, stack.head}
}

func (stack *Stack[T]) Pop() (T, error) {
	if stack.head == nil {
		return *new(T), errors.New("stack is empty")
	}

	oldHead := stack.head
	stack.head = stack.head.next

	return oldHead.value, nil
}
