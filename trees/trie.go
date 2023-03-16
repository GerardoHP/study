package trees

type Node struct {
	children map[int32]*Node
	isEnd    bool
}

type Trie struct {
	root *Node
}

func NewNode() *Node {
	return &Node{children: make(map[int32]*Node)}
}
func NewTrie() *Trie {
	return &Trie{root: NewNode()}
}

func (t *Trie) Insert(str string) {
	current := t.root
	for _, v := range str {
		i := GetIndex(v)
		if _, found := current.children[i]; !found {
			current.children[i] = NewNode()
		}

		current = current.children[i]
	}

	current.isEnd = true
}

func (t Trie) Search(str string) bool {
	current := t.root
	for _, v := range str {
		if _, found := current.children[v]; !found {
			return false
		}

		current = current.children[v]
	}

	return current.isEnd
}

func GetIndex(r int32) int32 {
	//return r - int32('A')
	return rgit
}
