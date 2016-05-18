package heap

import (
	"errors"
	"fmt"
)

type HeapItem interface {
	Priority() int
}

type Heap struct {
	data []HeapItem
}

func NewHeap(items []HeapItem) *Heap {
	h := Heap{data: items}
	h.init()
	return &h
}

func leftChild(i int) int {
	return i*2 + 1
}

func rightChild(i int) int {
	return i*2 + 2
}

func getParent(i int) int {
	return (i - 1) / 2
}

func swap(items []HeapItem, i int, j int) {
	tmp := items[i]
	items[i] = items[j]
	items[j] = tmp
}

func (self *Heap) init() {
	size := len(self.data)
	for i := getParent(size - 1); i >= 0; i-- {
		self.shiftDown(i)
	}
}

func (self *Heap) Push(item HeapItem) {
	self.data = append(self.data, item)
	size := len(self.data)
	self.shiftUp(size - 1)
}

func (self *Heap) Pop() (HeapItem, error) {
	if len(self.data) <= 0 {
		return nil, errors.New("pop empty heap")
	}
	size := len(self.data)
	res := self.data[0]
	self.data[0] = self.data[size-1]
	self.data = self.data[0 : size-1]
	if size-1 > 0 {
		self.shiftDown(0)
	}
	return res, nil
}

func (self *Heap) shiftUp(i int) {
	curVal := self.data[i]
	cur := i
	parent := getParent(cur)
	for {
		if cur <= 0 {
			break
		}
		if curVal.Priority() > self.data[parent].Priority() {
			self.data[cur] = self.data[parent]
			cur = parent
			parent = getParent(cur)
		} else {
			break
		}
	}
	self.data[cur] = curVal
}

func (self *Heap) shiftDown(i int) {
	curVal := self.data[i]
	size := len(self.data)
	cur := i
	left := leftChild(cur)
	for {
		if left >= size {
			break
		}
		if left+1 < size && self.data[left].Priority() < self.data[left+1].Priority() {
			left++
		}
		if self.data[left].Priority() > curVal.Priority() {
			self.data[cur] = self.data[left]
			cur = left
			left = leftChild(cur)
		} else {
			break
		}
	}
	self.data[cur] = curVal
}

func (self *Heap) toStr() string {
	return fmt.Sprintf("%+v", self.data)
}
