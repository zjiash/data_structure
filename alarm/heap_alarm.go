package alarm

import (
	"data_structure/heap"
	"errors"
	"fmt"
	"sync"
	"time"
)

type HeapAlarm struct {
	alarms      *heap.Heap
	mutex       *sync.Mutex
	milliSecond int
	enable      bool
	ticker      *time.Ticker
}

func NewHeapAlarm(milliSeconds ...int) *HeapAlarm {
	milliSecond := MilliSecondsPerTicker
	if len(milliSeconds) > 0 {
		milliSecond = milliSeconds[0]
	}
	return &HeapAlarm{
		alarms:      heap.NewEmptyHeap(),
		mutex:       &sync.Mutex{},
		milliSecond: milliSecond,
		enable:      false,
	}
}

func (self *HeapAlarm) start() {
	if self.enable {
		return
	}
	self.enable = true
	self.ticker = time.NewTicker(time.Millisecond * time.Duration(self.milliSecond))

	go func() {
		for t := range self.ticker.C {
			self.Check(t)
		}
	}()
}

func (self *HeapAlarm) stop() {
	if !self.enable {
		return
	}
	self.enable = false
	self.ticker.Stop()
}

func (self *HeapAlarm) Check(t time.Time) {
	fmt.Printf("check at %s\n", t)
	self.mutex.Lock()
	for {
		if self.alarms.IsEmpty() {
			self.stop()
			break
		}
		tmp, _ := self.alarms.Top()
		if alarm, ok := tmp.(*alarmItem); ok {
			now := time.Now()
			if alarm.When.Before(now) {
				self.alarms.Pop()
				if alarm.When.Before(now.Add(-Interval)) {
					panic(fmt.Sprintf("%s run at wrong time %s", alarm.toStr(), now))
				} else {
					go alarm.run()
				}
			} else {
				break
			}
		} else {
			self.mutex.Unlock()
			panic("must not appear")
		}
	}
	self.mutex.Unlock()
}

func (self *HeapAlarm) Add(when time.Time, call AlarmFunc, name string) error {
	now := time.Now()
	item := newAlarmItem(when, call, name)
	if when.Before(now) {
		return errors.New(fmt.Sprintf("error for add %s : alarm before now", item.toStr()))
	}

	self.mutex.Lock()
	err := self.alarms.Push(item)
	if err != nil {
		self.mutex.Unlock()
		return errors.New(fmt.Sprintf("error for add %s : can not add", item.toStr()))
	}
	self.start()
	self.mutex.Unlock()
	return nil
}

func (self *HeapAlarm) Cancel(name string) error {
	// todo
	return nil
}
