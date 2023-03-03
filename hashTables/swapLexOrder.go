package hashTables

import (
	"sort"
	"strings"
	"study/Interfaces"
)

type SwapLexOrder struct {
}

type sortRune []byte

func (s sortRune) Less(i, j int) bool {
	return s[i] < s[j]
}

func (s sortRune) Swap(i, j int) {
	s[i], s[j] = s[j], s[i]
}

func (s sortRune) Len() int {
	return len(s)
}

func (SwapLexOrder) Solution(str string, pairs [][]int) string {
	indexes := make(map[int][]int)
	for _, pair := range pairs {
		i := fixRange(pair[0])
		j := fixRange(pair[1])
		findAndAdd(indexes, i, j)
		findAndAdd(indexes, j, i)
	}

	var components [][]int
	for k, v := range indexes {
		arr := []int{k}
		arr = append(arr, v...)
		if len(components) == 0 {
			components = append(components, arr)
			continue
		}

		added := false
		for i := range components {
			if intersect, intersection := intersects(components[i], arr); intersect {
				components[i] = append(components[i], intersection...)
				added = true
				break
			}
		}

		if !added {
			components = append(components, arr)
		}
	}

	max := str
	for _, v := range components {
		order := orderKeys(v, str)
		max = createString(order, max)
	}

	return max
}

func unique(intSlice []int) []int {
	keys := make(map[int]bool)
	list := []int{}
	for _, entry := range intSlice {
		if _, value := keys[entry]; !value {
			keys[entry] = true
			list = append(list, entry)
		}
	}
	return list
}

func intersects(arr1, arr2 []int) (bool, []int) {
	h := make(map[int]bool)
	for _, v1 := range arr1 {
		if _, found := h[v1]; !found {
			h[v1] = true
		}
	}

	intersect := false
	var r []int
	for _, v := range arr2 {
		if _, found := h[v]; found {
			intersect = true
		} else {
			r = append(r, v)
		}
	}

	return intersect, r
}

func createString(m map[int]byte, str string) string {
	c := str
	sb := strings.Builder{}
	for i, v := range c {
		if s, found := m[i]; found {
			sb.WriteByte(s)
		} else {
			sb.WriteByte(byte(v))
		}
	}

	x := sb.String()
	return x
}

func orderKeys(component []int, str string) map[int]byte {
	var r []byte
	var n []int
	m := make(map[int]byte)
	for _, v := range component {
		r = append(r, str[v])
		n = append(n, v)
	}

	sort.Ints(n)
	sort.Sort(sort.Reverse(sortRune(r)))
	for i := range n {
		m[n[i]] = r[i]
	}

	return m
}

func findAndAdd(m map[int][]int, key int, value int) map[int][]int {
	if _, found := m[key]; !found {
		m[key] = []int{}
	}

	m[key] = append(m[key], value)
	return m
}

func (s SwapLexOrder) Execute() {
	panic("implement me")
}

func fixRange(i int) int {
	return i - 1
}

var _ Interfaces.Exercise = (*SwapLexOrder)(nil)
