package rbtree

import (
	"bytes"
	"fmt"
)

type RBTree struct {
	root *RBNode
}

func NewRBTree() *RBTree {
	return &RBTree{root: nil}
}

// 查询
func (self *RBTree) Search(data int64) bool {
	return self.searchNode(self.root, data)
}

func (self *RBTree) searchNode(node *RBNode, data int64) bool {
	if node == nil {
		return false
	}
	if node.value == data {
		return true
	} else if node.value > data {
		return self.searchNode(node.left, data)
	} else {
		return self.searchNode(node.right, data)
	}
	return false
}

// 插入
func (self *RBTree) Insert(data int64) {
	if self.root == nil {
		tmp := NewRBNode(data)
		tmp.color = BLACK
		self.root = tmp
	} else {
		self.insertNode(self.root, data)
	}
}

// 插入，BST的插入方法，插入成功之后进行调整
func (self *RBTree) insertNode(node *RBNode, data int64) {
	if node.value >= data {
		if node.left == nil {
			tmp := NewRBNode(data)
			tmp.parent = node
			node.left = tmp
			self.adjust(tmp)
		} else {
			self.insertNode(node.left, data)
		}
	} else {
		if node.right == nil {
			tmp := NewRBNode(data)
			tmp.parent = node
			node.right = tmp
			self.adjust(tmp)
		} else {
			self.insertNode(node.right, data)
		}
	}
}

func (self *RBTree) adjust(node *RBNode) {
	parent := node.parent
	if parent == nil {
		// 情形1：（递归后）新节点位于树的根上
		node.color = BLACK
		self.root = node
		return
	}

	if parent.color == BLACK {
		// 情形2：新节点的父节点P是黑色，性质不失效
		return
	} else {
		uncle := node.getUncle()
		grandParent := node.getGrandParent()
		if uncle != nil && uncle.color == RED {
			// 情形3：父节点和叔父节点均为红色
			// 处理：父节点和叔父节点转为黑色，祖父节点改成红色，递归处理祖父节点
			node.parent.color = BLACK
			uncle.color = BLACK
			grandParent.color = RED
			self.adjust(grandParent)
		} else {
			isLeft := (node == parent.left)
			isParentLeft := (node.parent == grandParent.left)
			if isLeft && isParentLeft {
				// 情形5：左左
				// 处理：祖父节点右旋
				node.parent.color = BLACK
				grandParent.color = RED
				self.rotateRight(grandParent)
			} else if !isLeft && isParentLeft {
				// 情形4：左右
				// 处理：父节点左旋，变成情形4左左
				self.rotateLeft(node.parent)
				// 现在的node.parent表示原来的祖父节点
				self.rotateRight(node.parent)

				node.color = BLACK
				node.left.color = RED
				node.right.color = RED
			} else if isLeft && !isParentLeft {
				// 情形4对称：右左
				self.rotateRight(node.parent)
				self.rotateLeft(node.parent)

				node.color = BLACK
				node.left.color = RED
				node.right.color = RED
			} else {
				// 情形5对称：右右
				node.parent.color = BLACK
				grandParent.color = RED
				self.rotateLeft(grandParent)
			}
		}
	}
}

func (self *RBTree) rotateLeft(node *RBNode) {
	changedRoot, err := node.leftRotate()
	if err != nil {
		fmt.Println(err.Error())
	} else {
		if changedRoot != nil {
			self.root = changedRoot
		}
	}
}

func (self *RBTree) rotateRight(node *RBNode) {
	changedRoot, err := node.rightRotate()
	if err != nil {
		fmt.Println(err.Error())
	} else {
		if changedRoot != nil {
			self.root = changedRoot
		}
	}
}

func inorderStr(node *RBNode) string {
	if node == nil {
		return ""
	}
	var buffer bytes.Buffer
	buffer.WriteString(inorderStr(node.left))
	buffer.WriteString(fmt.Sprintf(" %d", node.value))
	buffer.WriteString(inorderStr(node.right))
	return buffer.String()
}

func printTree(node *RBNode, prefix string, isRight bool) {
	if node != nil {
		var color string
		if node.color == RED {
			color = "RED"
		} else {
			color = "BLACK"
		}

		leftOrRight := ""
		if node.parent != nil {
			if node.parent.left == node {
				leftOrRight = "(l)"
			} else {
				leftOrRight = "(r)"
			}
		}

		newPrefix := prefix + "|   "
		if isRight {
			newPrefix = prefix + "    "
		}
		fmt.Printf("%s|-- %s%d(%s)\n", prefix, leftOrRight, node.value, color)
		printTree(node.left, newPrefix, false)
		printTree(node.right, newPrefix, true)
	}
}
