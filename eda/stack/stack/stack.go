package stack

import "errors"

type Stack struct {
	array []int // Slice de inteiros
}

func (s *Stack) Len() int {
	return len(s.array)
}

func (s *Stack) Push(value int) {
	s.array = append(s.array, value)
}

func (s *Stack) Pop() (int, error) {
	if len(s.array) <= 0 {
		return 0, errors.New("empty stack")
	}

	v := s.array[len(s.array)-1]
	s.array = s.array[:len(s.array)-1]

	return v, nil
}

func (s *Stack) Get(i int) (int, error) {
	if i >= len(s.array) {
		return 0, errors.New("out of bounds")
	}
	return s.array[i], nil
}

func Test(a, b int) int {
	return a + b
}
