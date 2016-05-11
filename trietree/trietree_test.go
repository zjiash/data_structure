package trietree

import (
	"testing"
)

func TestTrieTree(test *testing.T) {
	trietree := NewTrieTree()

	trietree.Insert("this")
	trietree.Insert("that")
	trietree.Insert("his")
	trietree.Insert("these")
	trietree.Insert("the")

	word := "the"
	if !trietree.Search(word) {
		test.Errorf("(%s) must be in trie tree\n", word)
	}

	word = "thes"
	if trietree.Search(word) {
		test.Errorf("(%s) must be not in trie tree\n", word)
	}
	if !trietree.SearchPrefix(word) {
		test.Errorf("(%s) must be prefix of trie tree\n", word)
	}

	PrintTree(trietree.root, "", -1)
}
