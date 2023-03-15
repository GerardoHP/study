package queue

type Queue []interface{}

func (q *Queue) Enqueue(i interface{}) {
	*q = append(*q, i)
}

func (q *Queue) Dequeue() interface{} {
	extract := (*q)[0]
	*q = (*q)[1:]
	return extract
}

func (q Queue) IsEmpty() bool {
	return len(q) == 0
}
