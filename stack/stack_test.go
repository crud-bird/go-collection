package stack_test

import (
	"go-collection/stack"
	"testing"
)

func TestSatck(t *testing.T) {
	s := stack.NewConcurrent()
	for i := 1; i < 10; i++ {
		s.Push(i)
	}

	t.Logf("%+v", s)

	t.Log(s.Pop())
	t.Log(s.Pop())

	t.Logf("%+v", s)
}
