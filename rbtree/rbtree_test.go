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

const TESTLEN int = 10000

func TestRBTree(test *testing.T) {
	rbtree := NewRBTree()

	rand.Seed(time.Now().UnixNano())
	intarr := rand.Perm(TESTLEN)
	value := int64(intarr[0])
	// fmt.Println(intarr)
	var buffer bytes.Buffer
	for index, data := range intarr {
		rbtree.Insert(int64(data))
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

	if !rbtree.Search(value) {
		test.Errorf("%d must be in rb tree\n", value)
	}

	if rbtree.Search(int64(TESTLEN)) {
		test.Errorf("%d must be not in rb tree\n", TESTLEN)
	}

}
