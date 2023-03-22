package queue

type Queue []interface{}

func (q *Queue) Enqueue(i interface{}) {
	*q = append(*q, i)
}

func NewQueueFromSlice(a []int) *Queue {
	var queue Queue
	for _, q := range a {
		queue = append(queue, q)
	}

	return &queue
}

func (q *Queue) Dequeue() interface{} {
	if len(*q) == 0 {
		return -1
	}
	
	extract := (*q)[0]
	*q = (*q)[1:]
	return extract
}

func (q Queue) IsNotEmpty() bool {
	return len(q) > 0
}

func (q Queue) IsEmpty() bool {
	return !q.IsNotEmpty()
}
