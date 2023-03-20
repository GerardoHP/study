package trees_test

import (
	"encoding/json"
	"io/ioutil"
	"os"
	"study/trees"
	"testing"
)

type entry struct {
	Input  DeleteFromBSTInput `json:"input"`
	Output *trees.Tree        `json:"output"`
}

type DeleteFromBSTInput struct {
	T       trees.Tree `json:"t"`
	Queries []int      `json:"queries"`
}

func TestDeleteFromBST_Test1(t *testing.T) {
	// arrange
	entry, err := newDeleteFromBstInputFromFile("testsFiles/deleteFromBST_test1.json")
	if err != nil {
		t.Fatal(err)
	}
	exercise := trees.DeleteFromBST{}
	input := entry.Input
	output := entry.Output

	// act
	result := exercise.Solution(&input.T, input.Queries)

	// assert
	if !compareTrees(result, output) {
		t.Fail()
	}
}

func TestDeleteFromBST_Test3(t *testing.T) {
	// arrange
	entry, err := newDeleteFromBstInputFromFile("testsFiles/deleteFromBST_test3.json")
	if err != nil {
		t.Fatal(err)
	}
	exercise := trees.DeleteFromBST{}
	input := entry.Input
	output := entry.Output

	// act
	result := exercise.Solution(&input.T, input.Queries)

	// assert
	if !compareTrees(result, output) {
		t.Fail()
	}
}

func TestDeleteFromBST_Test5(t *testing.T) {
	// arrange
	entry, err := newDeleteFromBstInputFromFile("testsFiles/deleteFromBST_test5.json")
	if err != nil {
		t.Fatal(err)
	}
	exercise := trees.DeleteFromBST{}
	input := entry.Input
	output := entry.Output

	// act
	result := exercise.Solution(&input.T, input.Queries)

	// assert
	if !compareTrees(result, output) {
		t.Fail()
	}
}

func TestDeleteFromBST_Test9(t *testing.T) {
	// arrange
	entry, err := newDeleteFromBstInputFromFile("testsFiles/deleteFromBST_test9.json")
	if err != nil {
		t.Fatal(err)
	}
	exercise := trees.DeleteFromBST{}
	input := entry.Input
	output := entry.Output

	// act
	result := exercise.Solution(&input.T, input.Queries)

	// assert
	if !compareTrees(result, output) {
		t.Fail()
	}
}

func compareTrees(t1, t2 *trees.Tree) bool {
	if t1 == nil || t2 == nil {
		if t1 == nil && t2 == nil {
			return true
		}

		return false
	}

	return t1.Value == t2.Value && compareTrees(t1.Left, t2.Left) && compareTrees(t1.Right, t2.Right)
}

func newDeleteFromBstInputFromFile(file string) (*entry, error) {
	//jsonFile, err := os.Open("files/test-9.json")
	jsonFile, err := os.Open(file)
	if err != nil {
		return nil, err
	}

	defer jsonFile.Close()
	byteValue, err := ioutil.ReadAll(jsonFile)
	if err != nil {
		return nil, err
	}

	var entry entry
	json.Unmarshal(byteValue, &entry)
	return &entry, nil
}
