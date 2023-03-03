package Trees

import (
	"fmt"
	"study/Interfaces"
)

type FindSubstrings struct{}

func (FindSubstrings) solution(words []string, parts []string) []string {
	trie := NewTrie()
	for _, str := range parts {
		trie.Insert(str)
	}

	ret := make([]string, 0)

	for _, word := range words {
		word = stringInTrie(word, trie)
		ret = append(ret, word)
	}

	return ret
}

func stringInTrie(str string, trie *Trie) string {
	maxIndex := -1
	maxSize := -1
	for i := range str {
		current := trie.root.children
		size := 0
		for current != nil {
			currentIndex := i + size
			if currentIndex >= len(str) {
				break
			}

			currentLetter := int32(str[currentIndex])
			if _, found := current[currentLetter]; !found {
				break
			}

			size++
			if current[currentLetter].isEnd {
				if size > maxSize {
					maxIndex = i
					maxSize = size
				}
			}

			current = current[currentLetter].children
		}
	}

	if maxSize > 0 {
		return GetString(str, maxIndex, maxSize)
	}

	return str
}

func GetString(str string, index, size int) string {
	arr3 := str[index : index+size]
	arr1 := str[:index]
	arr2 := str[index+size:]
	return fmt.Sprintf("%v[%v]%v", arr1, arr3, arr2)
}

func (f FindSubstrings) Execute() {
	words := []string{
		//"Apple",
		//"Melon",
		//"Orange",
		//"Watermelon"}
		"aaaaaaaaaaaaaaaaaaaaaaaaaaaaab",
		"aaaaaaaaaaaaaaaaaaaaaaaaaaaaab",
		"aaaaaaaaaaaaaaaaaaaaaaaaaaaaab",
		"aaaaaaaaaaaaaaaaaaaaaaaaaaaaab",
		"aaaaaaaaaaaaaaaaaaaaaaaaaaaaab",
		"aaaaaaaaaaaaaaaaaaaaaaaaaaaaa",
		"aaaaaaaaaaaaaaaaaaaaaaaaaaaaa",
		"aaaaaaaaaaaaaaaaaaaaaaaaaaaaaa",
		"aaaaaaaaaaaaaaaaaaaaaaaaaaaaaa",
		"aaaaaaaaaaaaaaaaaaaaaaaaaaaaac"}
	parts := []string{
		//"a",
		//"mel",
		//"lon",
		//"el",
		//"An"}
		"aaaaa",
		"bbbbb",
		"a",
		"aa",
		"aaaa",
		"AAAAA",
		"aaa",
		"aba",
		"aaaab",
		"c",
		"bbbb",
		"d",
		"g",
		"ccccc",
		"B",
		"C",
		"P",
		"D"}
	sol := f.solution(words, parts)
	for _, v := range sol {
		fmt.Println(v)
	}
}

var _ Interfaces.Exercise = (*FindSubstrings)(nil)
