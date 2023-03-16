package trees

import (
	"fmt"
	"study/interfaces"
)

type TreeToList struct {
}

var _ interfaces.Exercise = (*TreeToList)(nil)
var slice []interface{} = make([]interface{}, 1)

func (TreeToList) solution(t *Tree) *Tree {
	// return iterate(t)
	// return iterate(t)
	// return nil
	iterate(t)
	return nil
}

func iterate(t *Tree) {
	if t != nil {
		iterate(t.Left)
		//if len(slice) > 0 {
		//	aux := []int{t.Value.(int)}
		//	aux = append(aux, slice...)
		//	slice = aux
		//} else {
		//	slice[0] = t.Value.(int)
		//}
		slice = append(slice, t.Value)
		iterate(t.Right)
	}
}

func toTree(slice []interface{}) *Tree {
	//
	//for _, s := range slice {
	//
	//}
	return nil
}

//func iterate(t *Tree) *Tree {
//	if t != nil {
//		x := &Tree{}
//		x.Left = iterate(t.Left)
//		x.Value = t.Value
//		x.Right = iterate(t.Right)
//
//		return x
//		// iterate(t.Left)
//		// fmt.Println(t.Value)
//		// iterate(t.Right)
//		// return &Tree{Left: iterate(t.Left), Value: t.Value, Right: iterate(t.Right)}
//	}
//
//	return nil
//}

func (TreeToList) Execute() {
	x := TreeToList{}

	t1 := &Tree{Value: 10}
	t2 := &Tree{Value: 12}
	t3 := &Tree{Value: 15}
	t4 := &Tree{Value: 25}
	t5 := &Tree{Value: 30}
	t6 := &Tree{Value: 36}

	t3.Left = t6
	t2.Left = t4
	t2.Right = t5
	t1.Left = t2
	t1.Right = t3

	list := x.solution(t1)
	fmt.Println(list)
}
