package stack

import "fmt"


type Stack struct {
	elements []interface{}
}

func NewStack() *Stack {
	return &Stack{
		elements: make([]interface{}, 0),
	}
}

func (s *Stack) Push(data interface{}) {
	s.elements = append(s.elements, data)
}

func (s *Stack) Pop() interface{} {
	if len(s.elements) == 0 {
		return nil
	}

	n := len(s.elements) - 1
	data := s.elements[n]
	s.elements[n] = nil
	s.elements = s.elements[:n]
	return data
}

func (s *Stack) Peek() interface{} {
	if len(s.elements) == 0 {
		return nil
	}
	n := len(s.elements) - 1
	data := s.elements[n]
	return data
}

func (s *Stack) IsEmpty() bool {
	return len(s.elements) == 0
}

func (s *Stack) Size() int {
	return len(s.elements)
}

func (s *Stack) Print() {
	if s.IsEmpty() {
		fmt.Println("Stack is empty")
		return
	}
}
