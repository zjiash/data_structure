package trietree

const TREEWIDTH int = 26
const BASECHAR uint8 = 'a'

type TrieNode struct {
	hasWord  bool
	children [TREEWIDTH]*TrieNode
}

func NewTrieNode() *TrieNode {
	return &TrieNode{hasWord: false}
}
