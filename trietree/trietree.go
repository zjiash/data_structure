package trietree

import (
	"fmt"
)

type TrieTree struct {
	root *TrieNode
}

func NewTrieTree() *TrieTree {
	root := NewTrieNode()
	root.hasWord = true
	return &TrieTree{root: root}
}

func (self *TrieTree) Insert(word string) {
	node := self.root
	wordLen := len(word)
	for i := 0; i < wordLen; i++ {
		index := word[i] - BASECHAR
		if node.children[index] == nil {
			tmp := NewTrieNode()
			node.children[index] = tmp
		}
		node = node.children[index]
		if i+1 == wordLen {
			node.hasWord = true
		}
	}
}

func (self *TrieTree) find(word string) *TrieNode {
	node := self.root
	wordLen := len(word)
	for i := 0; i < wordLen && node != nil; i++ {
		index := word[i] - BASECHAR
		node = node.children[index]
	}
	return node
}

func (self *TrieTree) SearchPrefix(prefix string) bool {
	node := self.find(prefix)
	return node != nil
}

func (self *TrieTree) Search(word string) bool {
	node := self.find(word)
	if node == nil {
		return false
	}
	return node.hasWord
}

func PrintTree(node *TrieNode, prefix string, index int) {
	if node != nil {
		newPrefix := prefix + "|   "
		if index+1 >= TREEWIDTH {
			newPrefix = prefix + "    "
		}
		if index < 0 {
			fmt.Printf("%s|-- %s\n", prefix, "root")
		} else {
			value := int(BASECHAR) + index
			fmt.Printf("%s|-- %c[%d](%t)\n", prefix, value, index, node.hasWord)
		}
		for i := 0; i < TREEWIDTH; i++ {
			PrintTree(node.children[i], newPrefix, i)
		}
	}
}
