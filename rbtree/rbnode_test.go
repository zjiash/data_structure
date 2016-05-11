package rbtree

import (
	"fmt"
	"testing"
)

func addChild(value int64, parent *RBNode, isleft bool) *RBNode {
	son := NewRBNode(value)
	son.parent = parent
	if isleft {
		parent.left = son
	} else {
		parent.right = son
	}
	return son
}

func TestRBNodeRotate(test *testing.T) {
	root := NewRBNode(4)
	l := addChild(2, root, true)
	r := addChild(6, root, false)
	addChild(1, l, true)
	addChild(3, l, false)
	addChild(5, r, true)
	addChild(7, r, false)
	fmt.Println("original tree: ")
	printTree(root, "", true)

	newRoot, err := root.rightRotate()
	if err != nil || newRoot.value != 2 {
		test.Error("right rotate error")
	}
	fmt.Println("tree after right rotate 4: ")
	printTree(newRoot, "", true)

	newRoot2, err := newRoot.leftRotate()
	if err != nil || newRoot2.value != 4 {
		test.Error("left rotate error")
	}
	fmt.Println("tree after left rotate 2: ")
	printTree(newRoot2, "", true)

	newRoot3, err := newRoot2.left.rightRotate()
	if err != nil || newRoot3 != nil {
		test.Error("right rotate error")
	}
	fmt.Println("tree after right rotate 2: ")
	printTree(newRoot2, "", true)

	newRoot4, err := newRoot2.right.leftRotate()
	if err != nil || newRoot4 != nil {
		test.Error("left rotate error")
	}
	fmt.Println("tree after left rotate 6: ")
	printTree(newRoot2, "", true)
}
