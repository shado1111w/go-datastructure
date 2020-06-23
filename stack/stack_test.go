package stack

import (
	"testing"
)

type Stack interface {
	Init() bool
	Push(e interface{}) bool
	Pop() (e interface{}, ok bool)
	GetTop() (e interface{}, ok bool)
}

func TestStack(t *testing.T) {
	var stack Stack
	stack = new(SqStack)
	// stack = new(StackNode)
	// stack = new(IntStack)

	ok := stack.Init()
	t.Logf("init ok?%t\n", ok)

	ok = stack.Push(1)
	t.Logf("push ok?%t\n", ok)
	ok = stack.Push(3)
	t.Logf("push ok?%t\n", ok)
	ok = stack.Push(5)
	t.Logf("push ok?%t\n", ok)

	e, ok := stack.Pop()
	t.Logf("pop ok?%t e:%d\n", ok, e)
	e, ok = stack.Pop()
	t.Logf("pop ok?%t e:%d\n", ok, e)

	e, ok = stack.GetTop()
	t.Logf("pop ok?%t e:%d\n", ok, e)
}
