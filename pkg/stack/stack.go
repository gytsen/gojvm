package stack

import (
	log "github.com/sirupsen/logrus"
)

const (
	InitialStackSize = 1000
)

// Word represents an item on the stack
type Word int32

// The Stack struct represents the stack itself
type Stack struct {
	stack    []Word
	capacity int
}

// New creates a new Stack struct with a fixed initial capacity
func New() *Stack {

	stack := make([]Word, 0, InitialStackSize)

	return &Stack{
		stack:    stack,
		capacity: InitialStackSize,
	}
}

// Push a new word on the stack. If the stack is of insufficient size
// the stack capacity gets doubled automatically
func (s *Stack) Push(w Word) {
	if len(s.stack) >= s.capacity {
		s.grow()
	}

	s.stack = append(s.stack, w)
}

// Pop a word from the stack. If the stack is empty, the function panics.
func (s *Stack) Pop() Word {
	length := len(s.stack)

	if length == 0 {
		log.Panic("Attempted to pop from an empty stack")
	}

	r := s.stack[length-1]
	s.stack = s.stack[:length-1]

	return r
}

// Top returns the top of the stack. If the stack is empty, function panics
func (s *Stack) Top() Word {
	length := len(s.stack)

	if length == 0 {
		log.Panic("Attempted to top from an empty stack")
	}

	return s.stack[length-1]
}

// Size returns the current size of the stack
func (s *Stack) Size() int {
	return len(s.stack)
}

// double the stack in capacity. only called internally
func (s *Stack) grow() {
	size := s.Size()
	newCapacity := s.capacity * 2

	newStack := make([]Word, size, newCapacity)
	copy(newStack, s.stack)
	s.stack = newStack
}
