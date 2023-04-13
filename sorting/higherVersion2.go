package sorting

import "study/interfaces"

type HigherVersion2 struct{}

func (h HigherVersion2) Execute() {
	//TODO implement me
	panic("implement me")
}

var _ interfaces.Exercise = (*HigherVersion2)(nil)

func (HigherVersion2) Solution(ver1 string, ver2 string) int {
	return 0
}
