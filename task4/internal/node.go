package internal

type node[T any] struct {
	Value T
	Next *node[T]
}