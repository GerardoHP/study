package stack

type Stack []interface{}

func (s *Stack) Push(i interface{}) {
	*s = append(*s, i)
}

func (s *Stack) Pop() interface{} {
	l := len(*s) - 1
	if l == -1 {
		return ""
	}

	extract := (*s)[l]
	*s = (*s)[:l]
	return extract
}
