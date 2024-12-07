package set

import "slices"

type Set[T comparable] struct {
	set []T
}

func (s *Set[T]) Append(e T) bool {
	if slices.Contains(s.set, e) {
		return false
	}

	s.set = append(s.set, e)

	return true
}

func (s *Set[T]) Len() int {
	return len(s.set)
}

func (s *Set[T]) Contains(e T) bool {
	return slices.Contains(s.set, e)
}

func (s *Set[T]) Reverse() {
	slices.Reverse(s.set)
	return
}
