package core

import "errors"

/*
Stack data structure implementation
*/
type Stack[T any] struct {
	slice []T
}

func NewStack[T any]() *Stack[T] {
	slice := make([]T, 0)
	return &Stack[T]{slice: slice}
}

func (s *Stack[T]) Push(el T) {
	s.slice = append(s.slice, el)
}

func (s *Stack[T]) deleteLast() {
	if len(s.slice) > 0 {
		s.slice = s.slice[:len(s.slice)-1]
	}
}

func (s *Stack[T]) Pop() (T, error) {
	defer s.deleteLast()
	if len(s.slice) > 0 {
		return s.slice[len(s.slice)-1:][0], nil
	} else {
		var zeroVal T
		return zeroVal, errors.New("Stack is empty")
	}
}

func (s *Stack[T]) Size() int {
	return len(s.slice)
}
