package lists

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
)

type entry struct {
	Input input `json:"input"`
}

type input struct {
	L []int `json:"l"`
}

type ListNode struct {
	Value interface{}
	Next  *ListNode
}

func NewListNodeFromValue(d int) *ListNode {
	return &ListNode{Value: d}
}

func NewListNodeFromSlice(arr []int) *ListNode {
	if len(arr) == 0 {
		return nil
	}

	rln := NewListNodeFromValue(arr[0])
	for _, v := range arr[1:] {
		rln.Append(v)
	}

	return rln
}

func NewListFromFile() (*ListNode, error) {
	jsonFile, err := os.Open("files/test-21.json")
	if err != nil {
		return nil, err
	}

	defer jsonFile.Close()
	byteValue, _ := ioutil.ReadAll(jsonFile)
	var entry entry
	json.Unmarshal(byteValue, &entry)
	return NewListNodeFromSlice(entry.Input.L), nil
}

func (ln *ListNode) Append(d int) {
	end := NewListNodeFromValue(d)
	n := ln
	for n.Next != nil {
		n = n.Next
	}

	n.Next = end
}

func (ln *ListNode) Print() {
	for ln != nil {
		fmt.Printf(" %v,", ln.Value)
		ln = ln.Next
	}

	fmt.Println()
}
