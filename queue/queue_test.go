package queue_test

import (
	"study/queue"
	"testing"
)

func TestQueue_Dequeue(t *testing.T) {
	// arrange
	q := queue.Queue{}
	q.Enqueue(100)
	q.Enqueue(200)
	q.Enqueue(300)

	// act
	dequeued := q.Dequeue()

	// assert
	if dequeued != 100 {
		t.Fail()
	}
}
