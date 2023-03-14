package queue

type Queue []interface{}

func (q *Queue) Enqueue(i int) {
	*q = append(*q, i)
}

func (q *Queue) Dequeue() interface{} {
	extract := (*q)[0]
	*q = (*q)[1:]
	return extract
}
