package rbtree

import (
	"bytes"
	"fmt"
	"math/rand"
	"sort"
	"strings"
	"testing"
	"time"
)

func TestRBTree(test *testing.T) {
	rbtree := NewRBTree()

	rand.Seed(time.Now().UnixNano())
	intarr := rand.Perm(10000)
	// fmt.Println(intarr)
	var buffer bytes.Buffer
	for index, data := range intarr {
		rbtree.insert(int64(data))
		buffer.WriteString(fmt.Sprintf(" %d", index))
	}
	sort.Ints(intarr)
	// fmt.Println(intarr)
	orderedStr := strings.TrimSpace(inorderStr(rbtree.root))
	// fmt.Println("inorder string: " + orderedStr)
	if orderedStr != strings.TrimSpace(buffer.String()) {
		test.Errorf("inorderStr error: %s \n", orderedStr)
	}
	// fmt.Println("print tree")
	// printTree(rbtree.root, "", true)
}
