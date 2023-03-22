package trees

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
)

type Tree struct {
	Value interface {
	} `json:"value, int"`
	Left  *Tree `json:"left"`
	Right *Tree `json:"right"`
}

type entry struct {
	Input input `json:"input"`
}

type input struct {
	T1 Tree `json:"t1"`
	T2 Tree `json:"t2"`
}

func NewTreesFromFile() (*Tree, *Tree, error) {
	jsonFile, err := os.Open("files/test-9.json")
	if err != nil {
		return nil, nil, err
	}

	defer jsonFile.Close()
	bytevalue, _ := ioutil.ReadAll(jsonFile)
	var entry entry
	json.Unmarshal(bytevalue, &entry)
	return &entry.Input.T1, &entry.Input.T2, nil
}

func (t *Tree) InOrderTraversal() {
	if t != nil {
		t.Left.InOrderTraversal()
		fmt.Printf("%v, ", t.Value)
		t.Right.InOrderTraversal()
	}
}

func (t *Tree) PreOrderTraversal() {
	if t != nil {
		fmt.Println(t.Value)
		t.Left.PreOrderTraversal()
		t.Right.PreOrderTraversal()
	}
}

func (t *Tree) PostOrderTraversal() {
	if t != nil {
		t.Left.PostOrderTraversal()
		t.Right.PostOrderTraversal()
		fmt.Println(t.Value)
	}
}

func (t *Tree) LevelOrderTraversal() {
	if t == nil {
		return
	}

	queue := make([]*Tree, 0)
	queue = append(queue, t)

	for len(queue) != 0 {
		node := queue[0]
		queue = queue[1:]

		fmt.Println(node.Value)
		if node.Left != nil {
			queue = append(queue, node.Left)
		}

		if node.Right != nil {
			queue = append(queue, node.Right)
		}
	}
}
