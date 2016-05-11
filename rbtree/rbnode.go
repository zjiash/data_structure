package rbtree

import (
	"errors"
)

const (
	RED   bool = true
	BLACK bool = false
)

type RBNode struct {
	value               int64
	color               bool
	left, right, parent *RBNode
}

func (self *RBNode) leftRotate() (*RBNode, error) {
	var root *RBNode
	if self == nil {
		return root, nil
	}

	parent := self.parent
	rightChild := self.right
	if rightChild == nil {
		return root, errors.New("left rotate node without right child")
	}

	// rotate
	rightChildLeft := rightChild.left
	rightChild.left = self
	self.parent = rightChild
	self.right = rightChildLeft
	if rightChildLeft != nil {
		rightChildLeft.parent = self
	}

	if parent == nil {
		root = rightChild
	} else {
		// 判断左右
		if parent.left == self {
			parent.left = rightChild
		} else {
			parent.right = rightChild
		}
	}
	rightChild.parent = parent

	return root, nil
}

func (self *RBNode) rightRotate() (*RBNode, error) {
	var root *RBNode
	if self == nil {
		return root, nil
	}

	leftChild := self.left
	parent := self.parent
	if leftChild == nil {
		return root, errors.New("right rotate node without left child")
	}

	// rotate
	leftChildRight := leftChild.right
	leftChild.right = self
	self.parent = leftChild
	self.left = leftChildRight
	if leftChildRight != nil {
		leftChildRight.parent = self
	}

	if parent == nil {
		root = leftChild
	} else {
		// 判断左右
		if parent.left == self {
			parent.left = leftChild
		} else {
			parent.right = leftChild
		}
	}
	leftChild.parent = parent
	return root, nil
}

func (self *RBNode) getGrandParent() *RBNode {
	parent := self.parent
	if parent == nil {
		return nil
	}

	return parent.parent
}

func (self *RBNode) getUncle() *RBNode {
	grandParent := self.getGrandParent()
	if grandParent == nil {
		return nil
	}

	if self.parent == grandParent.left {
		return grandParent.right
	} else {
		return grandParent.left
	}
}

func NewRBNode(value int64) *RBNode {
	return &RBNode{
		value: value,
		color: RED,
	}
}
