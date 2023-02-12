package LinkedList

type Node struct {
	Next *Node
	Data int
}

func NewNode(d int) *Node {
	return &Node{Data: d}
}

func (node *Node) appendToTail(d int) {
	end := NewNode(d)
	n := node
	for n.Next != nil {
		n = n.Next
	}

	n.Next = end
}
