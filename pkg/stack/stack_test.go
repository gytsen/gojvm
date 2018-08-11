package stack

import (
	"testing"
)

func TestNew(t *testing.T) {
	s := New()
	size := s.Size()
	if size != 0 {
		t.Errorf("expected size of new stack to be 0, not %d", size)
	}

	capacity := cap(s.stack)
	if capacity != InitialStackSize {
		t.Errorf("expected capacity of new stack to be %d, not %d", InitialStackSize, capacity)
	}
}

func TestPush(t *testing.T) {
	s := New()
	data := []Word{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	for index, n := range data {
		s.Push(n)
		if s.stack[index] != data[index] {
			t.Errorf("pushed stack value is %d instead of %d", s.stack[index], data[index])
		}
	}
}

func TestTop(t *testing.T) {
	s := New()
	data := []Word{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	for index, n := range data {
		s.Push(n)
		if s.stack[index] != data[index] {
			t.Errorf("pushed stack value is %d instead of %d", s.stack[index], data[index])
		}

		top := s.Top()
		if top != data[index] {
			t.Errorf("Top returned %d instead of %d", top, data[index])
		}
	}
}

func TestPop(t *testing.T) {
	s := New()
	data := []Word{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	for _, n := range data {
		s.Push(n)
	}

	for i := len(data); i > 0; i-- {
		if i != len(s.stack) {
			t.Errorf("Expected stack to be length %d before pop, not %d", i, len(s.stack))
		}

		p := s.Pop()

		index := i - 1
		length := len(s.stack)

		if p != data[index] {
			t.Errorf("Pop value returned %d instead of %d", p, data[index])
		}

		if index != length {
			t.Errorf("Expected stack to be length %d after pop, not %d", index, length)
		}
	}
}

func TestGrow(t *testing.T) {
	s := New()
	data := []Word{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}

	for _, n := range data {
		s.Push(n)
	}

	oldLength := len(s.stack)
	oldCap := cap(s.stack)

	s.grow()

	newLength := len(s.stack)
	newCap := cap(s.stack)

	if oldLength != newLength {
		t.Errorf("expected old stack length to match new length, instead got %d:%d", oldLength, newLength)
	}

	if (2 * oldCap) != newCap {
		t.Errorf("Expected new cap to be 2 times old cap. Instead got old:%d, new:%d", oldCap, newCap)
	}

	for i, n := range data {
		if s.stack[i] != n {
			t.Errorf("Mismatched stack contents after growing, expected %d, not %d", n, s.stack[i])
		}
	}
}

func TestPopPanic(t *testing.T) {
	defer func() {
		if r := recover(); r == nil {
			t.Errorf("Pop on an empty stack did not panic")
		}
	}()

	s := New()
	s.Pop()
}

func TestTopPanic(t *testing.T) {
	defer func() {
		if r := recover(); r == nil {
			t.Errorf("Top on an empty stack did not panic")
		}
	}()

	s := New()
	s.Top()
}
