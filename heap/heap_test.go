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

func (self Integer) Priority() int64 {
	return int64(self)
}

func TestHeap(test *testing.T) {
	testData := make([]interface{}, 8)
	testData[0] = Integer(1)
	testData[1] = Integer(2)
	testData[2] = Integer(3)
	testData[3] = Integer(4)
	testData[4] = Integer(5)
	testData[5] = Integer(6)
	testData[6] = Integer(7)
	testData[7] = Integer(8)
	ht, _ := NewHeap(testData)
	for {
		if ht.IsEmpty() {
			break
		}
		tmp, _ := ht.Pop()
		if value, ok := tmp.(Integer); ok {
			fmt.Printf("%d ", value.Priority())
		} else {
			test.Error("pop type error")
		}
	}

	rand.Seed(time.Now().UnixNano())
	intarr := rand.Perm(TESTLEN)
	// fmt.Println(intarr)
	h := NewEmptyHeap()
	for index := range intarr {
		h.Push(Integer(intarr[index]))
	}
	sort.Ints(intarr)
	for index := range intarr {
		tmp, err := h.Pop()
		if err != nil {
			test.Error("pop error")
		} else if value, ok := tmp.(Integer); ok {
			if int(value.Priority()) != intarr[TESTLEN-index-1] {
				test.Errorf("heap sort error: [%d] %d != %d\n", index, value.Priority(), intarr[TESTLEN-index-1])
			}
		} else {
			test.Error("pop type error")
		}
	}

	if !h.IsEmpty() {
		test.Error("heap must be empty now!")
	}
}
