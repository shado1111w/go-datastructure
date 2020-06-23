package queue

import (
	"testing"
)

type Queue interface {
	Init() bool
	EnQueue(e int) bool
	DeQueue() (int, bool)
	GetHead() (int, bool)
}

func TestQueue(t *testing.T) {
	var queue Queue
	// queue = new(SqQueue)
	queue = new(LinkQueue)

	ok := queue.Init()
	t.Logf("init ok?%t\n", ok)

	for i := 1; i <= 5; i++ {
		ok = queue.EnQueue(i)
		t.Logf("enqueue ok?%t\n", ok)
	}

	for i := 0; i < 5; i++ {
		e, ok := queue.DeQueue()
		if ok {
			t.Logf("dequeue e:%d", e)
		}
	}

	for i := 6; i <= 10; i++ {
		ok = queue.EnQueue(i)
		t.Logf("enqueue ok?%t\n", ok)
	}

	for i := 0; i < 5; i++ {
		e, ok := queue.DeQueue()
		if ok {
			t.Logf("dequeue e:%d", e)
		}
	}

}
