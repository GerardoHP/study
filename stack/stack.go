package stack

type Stack []interface{}

func NewStackFromSlice(a []int) *Stack {
	var stk Stack
	for _, v := range a {
		stk = append(stk, v)
	}

	return &stk
}

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

func (s Stack) IsNotEmpty() bool {
	return len(s) > 0
}

func (s Stack) IsEmpty() bool {
	return !s.IsNotEmpty()
}
