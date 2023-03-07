package hashTables

import (
	"sort"
	"strings"
	"study/Interfaces"
)

type component []int

type SwapLexOrder struct {
}

var (
	keys             []int
	current          int
	currentComponent *component
	components       []*component
	queue            []int
	visited          map[int]bool
)

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

func restart() {
	keys = []int{}
	current = -1
	currentComponent = nil
	components = []*component{}
	queue = []int{}
	visited = make(map[int]bool)
}

func (s *SwapLexOrder) Solution(str string, pairs [][]int) string {
	restart()
	edges := make(map[int][]int)
	for _, pair := range pairs {
		i := fixRange(pair[0])
		j := fixRange(pair[1])
		findAndAdd(edges, i, j)
		findAndAdd(edges, j, i)
	}

	visited = make(map[int]bool)
	keys = getKeys(edges)
	nextRound(edges)
	visited[current] = true
	//for i := 0; i < len(keys); i++ {
	for {
		for len(queue) > 0 {
			v, q := shift(queue)
			queue = q
			queueUp(edges, v)
			currentComponent.append(v)
			keys = removeKey(keys, v)
		}

		if len(keys) == 0 && len(queue) == 0 {
			break
		}

		nextRound(edges)
	}

	max := str
	for _, v := range components {
		order := orderKeys(*v, str)
		max = createString(order, max)
	}

	return max
}

func nextRound(edges map[int][]int) {
	if len(keys) == 0 {
		current = -1
		return
	}

	current, keys = shift(keys)
	currentComponent = newComponent(current)
	components = append(components, currentComponent)
	queueUp(edges, current)
}

func newComponent(i int) *component {
	n := []int{i}
	c := component(n)
	return &c
}

func shift(slice []int) (int, []int) {
	first := slice[0]
	slice = slice[1:]
	return first, slice
}

func queueUp(edges map[int][]int, key int) {
	if values, found := edges[key]; found {
		for _, value := range values {
			if _, f := visited[value]; !f {
				queue = append(queue, value)
				visited[value] = true
			}
		}
	}
}

func removeKey(keys []int, key int) []int {
	var r []int
	for _, k := range keys {
		if k != key {
			r = append(r, k)
		}
	}

	return r
}

func getKeys(edges map[int][]int) (keys []int) {
	for k := range edges {
		keys = append(keys, k)
	}

	return
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

func (c *component) append(v int) {
	exist := false
	for _, val := range *c {
		if v == val {
			exist = true
			break
		}
	}

	if !exist {
		*c = append(*c, v)
	}
}

func (s SwapLexOrder) Execute() {
	panic("implement me")
}

func fixRange(i int) int {
	return i - 1
}

var _ Interfaces.Exercise = (*SwapLexOrder)(nil)
