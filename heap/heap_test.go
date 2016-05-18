package heap

import (
	"fmt"
	"math/rand"
	"sort"
	"testing"
	"time"
)

const TESTLEN int = 10000

type Integer int

func (self Integer) Priority() int {
	return int(self)
}

func TestHeap(test *testing.T) {
	testData := make([]HeapItem, 8)
	testData[0] = Integer(1)
	testData[1] = Integer(2)
	testData[2] = Integer(3)
	testData[3] = Integer(4)
	testData[4] = Integer(5)
	testData[5] = Integer(6)
	testData[6] = Integer(7)
	testData[7] = Integer(8)
	ht := NewHeap(testData)
	for {
		if tmp, err := ht.Pop(); err == nil {
			fmt.Printf("%d ", tmp.Priority())
		} else {
			fmt.Printf("\n")
			break
		}
	}

	rand.Seed(time.Now().UnixNano())
	intarr := rand.Perm(TESTLEN)
	// fmt.Println(intarr)
	h := NewHeap([]HeapItem{})
	for index := range intarr {
		h.Push(Integer(intarr[index]))
	}
	sort.Ints(intarr)
	for index := range intarr {
		if tmp, err := h.Pop(); err != nil || tmp.Priority() != intarr[TESTLEN-index-1] {
			test.Errorf("heap sort error: [%d] %d != %d\n", index, tmp.Priority(), intarr[TESTLEN-index-1])
		}
	}
}
