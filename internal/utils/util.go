package util 

import "fmt"

type Set[T comparable] struct{
	elements map[T]struct{}
}

func NewSet[T comparable]() *Set[T] {
    return &Set[T]{elements: make(map[T]struct{})}
}

func (s *Set[T]) Add(value T) {
    s.elements[value] = struct{}{}
}

func (s *Set[T]) Remove(value T) {
    delete(s.elements, value)
}

func (s *Set[T]) Contains(value T) bool {
    _, exists := s.elements[value]
    return exists
}

func (s *Set[T]) Size() int {
    return len(s.elements)
}

// Print all elements in the set
func (s *Set[T]) Print() {
    for key := range s.elements {
        fmt.Println(key)
    }
}
