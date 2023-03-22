package trees

import (
	"strings"
	"study/interfaces"
	"study/utils"
)

type LongestPath struct{}

func (l LongestPath) Execute() {
	//TODO implement me
	panic("implement me")
}

var _ interfaces.Exercise = (*LongestPath)(nil)

func (LongestPath) Solution(fileSystem string) int {
	if len(fileSystem) == 0 {
		return 0
	}

	ar := strings.Split(fileSystem, "\f")
	maxLength := 0
	count := make(map[int]int)
	for _, s := range ar {
		tabIndex := strings.LastIndex(s, "\t")
		count[tabIndex+1] = len(s[tabIndex+1:])

		if strings.Index(s, ".") != -1 {
			length := 0
			for i := 0; i <= tabIndex+1; i++ {
				length += count[i]
				length++
			}

			length--
			maxLength = utils.MaxInt(length, maxLength)
		}
	}

	return maxLength
}
