package trees

import (
	"fmt"
	"study/interfaces"
)

type FindProfession struct {
}

type profession string

const (
	engineer profession = "Engineer"
	doctor   profession = "Doctor"
)

var _ interfaces.Exercise = (*FindProfession)(nil)

func (FindProfession) solution(level int, pos int) string {
	return string(getPosition(level, pos))
}

func getPosition(level int, pos int) profession {
	completeTree := createTree(engineer, level, 0)
	//return LevelOrderTraversal(completeTree, pos)
	return PreOrderTraversal(completeTree, pos, 0)
}

func PreOrderTraversal(t *Tree, pos int, currentPos int) profession {
	if t != nil {
		if currentPos == pos {
			return t.Value.(profession)
		}

		return PreOrderTraversal(t.Left, pos, currentPos+1)
		return PreOrderTraversal(t.Right, pos, currentPos+1)
	}

	return engineer
}

func LevelOrderTraversal(t *Tree, position int) profession {
	queue := make([]*Tree, 0)
	queue = append(queue, t)
	i := 0
	for len(queue) != 0 {
		node := queue[0]
		queue = queue[1:]

		if node.isLeaf() {
			i++
		}

		if i == position {
			return node.Value.(profession)
		}

		if node.Left != nil {
			queue = append(queue, node.Left)
		}

		if node.Right != nil {
			queue = append(queue, node.Right)
		}
	}

	return engineer
}

func createTree(father profession, level int, currentLevel int) *Tree {
	if level == currentLevel {
		return nil
	}

	t := &Tree{}
	if father == engineer {
		t.Value = engineer
		t.Left = createTree(engineer, level, currentLevel+1)
		t.Right = createTree(doctor, level, currentLevel+1)
	} else {
		t.Value = doctor
		t.Left = createTree(doctor, level, currentLevel+1)
		t.Right = createTree(engineer, level, currentLevel+1)
	}

	return t
}

func (FindProfession) Execute() {
	e := FindProfession{}
	fmt.Printf("el resultado es %v \n", e.solution(4, 3))
}
