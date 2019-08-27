package queue_test

import (
	"algorithm/structure/queue"
	"testing"
)

func TestQueue(t *testing.T) {
	q := queue.NewConcurrent()
	t.Log(q.IsEmpty())
	for i := 1; i < 20; i++ {
		q.Enqueue(i)
	}

	t.Log(q.Dequeue())
	t.Log(q.Dequeue())
	t.Log(q)
	for i := 100; i < 115; i++ {
		q.Enqueue(i)
	}

	t.Log(q.IsFull())
	t.Logf("%+v", q)
}
