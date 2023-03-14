package stack

import (
	"strings"
	"study/Interfaces"
)

type SimplifyPath struct {
}

func (s SimplifyPath) Execute() {
	//TODO implement me
	panic("implement me")
}

var _ Interfaces.Exercise = (*SimplifyPath)(nil)

func (s SimplifyPath) Solution(path string) string {
	strs := strings.Split(path, "/")
	stk := Stack{}
	for _, str := range strs {
		switch str {
		case "..":
			stk.Pop()
		case ".":
		case "":
		default:
			stk.Push(str)
		}
	}

	str := stk.toString()
	return str
}

func (stk Stack) toString() string {
	if len(stk) == 0 {
		return "/"
	}
	strBuilder := strings.Builder{}
	for _, str := range stk {
		strBuilder.WriteString("/")
		strBuilder.WriteString(str.(string))
	}

	return strBuilder.String()
}
