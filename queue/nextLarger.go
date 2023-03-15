package queue

import "study/Interfaces"

type NextLarger struct {
}

func (n NextLarger) Execute() {
	//TODO implement me
	panic("implement me")
}

var _ Interfaces.Exercise = (*NextLarger)(nil)

func (NextLarger) Solution(a []int) []int {
	sln := solution(a)
	return sln
}

func solution(a []int) []int {
	var r []int
	for i, v := range a {
		queue := sliceToQueue(a[i:])
		val := -1
		for !queue.IsEmpty() {
			q := queue.Dequeue().(int)
			if q > v {
				val = q
				break
			}
		}

		r = append(r, val)
	}

	return r
}

func sliceToQueue(slc []int) Queue {
	q := Queue{}
	for _, s := range slc {
		q.Enqueue(s)
	}

	return q
}
