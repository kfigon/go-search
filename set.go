package main

type Set[T comparable] map[T]struct{}

func NewSet[T comparable]() *Set[T] {
	return &Set[T]{}
}

func(s *Set[T]) Add(v T) {
	(*s)[v] = struct{}{}
}

func(s *Set[T]) Contains(v T) bool {
	_, ok := (*s)[v]
	return ok
}

func(s *Set[T]) Len() int {
	return len(*s)
}