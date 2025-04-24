package core

import (
	"errors"
	"sync"
)

/*
Stack data structure implementation
*/
type Stack[T any] struct {
	slice []T
	mtx   sync.RWMutex
}

func NewStack[T any]() *Stack[T] {
	slice := make([]T, 0)
	return &Stack[T]{slice: slice}
}

func (s *Stack[T]) Push(el T) {
	s.slice = append(s.slice, el)
}

func (s *Stack[T]) SafePush(el T) {
	s.mtx.Lock()
	s.slice = append(s.slice, el)
	s.mtx.Unlock()
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
		return zeroVal, errors.New("stack is empty")
	}
}

func (s *Stack[T]) Size() int {
	return len(s.slice)
}
