package stack

import (
	"fmt"
	"strconv"
	"strings"
	"study/Interfaces"
)

type DecodeString struct {
}

var _ Interfaces.Exercise = (*DecodeString)(nil)

func (s DecodeString) Execute() {
	//TODO implement me
	panic("implement me")
}

func (DecodeString) Solution(s string) string {
	str, _ := decode(s)
	return str
}

func decode(s string) (string, int) {
	stk := Stack{}
	i := 0
	for ; i < len(s); i++ {
		str := s[i]
		switch str {
		case '[':
			multiplier, _ := strconv.Atoi(stk.Pop().(string))
			newString, j := decode(s[i+1:])
			i = i + j
			stk.Push(multiplyString(newString, multiplier))
		case ']':
			return stk.stackToString(), i + 1
		case '0', '1', '2', '3', '4', '5', '6', '7', '8', '9':
			aux := stk.Pop().(string)
			_, err := strconv.Atoi(aux)
			if err != nil {
				stk.Push(aux)
				stk.Push(string(str))
			} else {
				n := fmt.Sprintf("%v%v", aux, string(str))
				stk.Push(n)
			}
		default:
			stk.Push(string(str))
		}
	}

	return stk.stackToString(), i
}

func multiplyString(str string, n int) string {
	s := strings.Builder{}
	i := 0
	for i < n {
		s.WriteString(str)
		i++
	}

	return s.String()
}

func (stk Stack) stackToString() string {
	strBuilder := strings.Builder{}
	for _, str := range stk {
		strBuilder.WriteString(str.(string))
	}

	return strBuilder.String()
}
