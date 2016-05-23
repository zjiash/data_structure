package alarm

import (
	"errors"
	"fmt"
	"sync"
	"time"
)

const (
	Second string = "Second"
	Minute        = "Minute"
	Hour          = "Hour"
)

type alarmListNode struct {
	item        *alarmItem
	TriggerSec  int
	TriggerMin  int
	TriggerHour int
	next        *alarmListNode
}

func newAlarmListNode(when time.Time, call AlarmFunc, name string) (*alarmListNode, error) {
	now := time.Now()
	innerItem := newAlarmItem(when, call, name)
	if when.Before(now) {
		return nil, errors.New(fmt.Sprintf("error for add %s : alarm before now", innerItem.toStr()))
	}

	diff := when.Sub(now)
	item := &alarmListNode{
		item:        innerItem,
		TriggerSec:  int(diff.Seconds()) % 60,
		TriggerMin:  int(diff.Minutes()) % 60,
		TriggerHour: int(diff.Hours()),
	}
	return item, nil
}

type TimerWheelAlarm struct {
	secList     [60]*alarmListNode
	minList     [60]*alarmListNode
	hourList    [24]*alarmListNode
	curSec      int
	curMin      int
	curHour     int
	mutex       *sync.Mutex
	milliSecond int
	ticker      *time.Ticker
}

func NewTimerWheelAlarm() *TimerWheelAlarm {
	alarm := &TimerWheelAlarm{
		mutex:       &sync.Mutex{},
		milliSecond: 1000,
	}
	alarm.start()
	return alarm
}

func (self *TimerWheelAlarm) start() {
	self.ticker = time.NewTicker(time.Millisecond * time.Duration(self.milliSecond))

	go func() {
		for t := range self.ticker.C {
			self.Check(t)
		}
	}()
}

func (self *TimerWheelAlarm) addToList(item *alarmListNode, lType string) {
	switch lType {
	case Second:
		item.next = self.secList[item.TriggerSec]
		self.secList[item.TriggerSec] = item
	case Minute:
		item.next = self.minList[item.TriggerMin]
		self.minList[item.TriggerMin] = item
	case Hour:
		item.next = self.hourList[item.TriggerHour]
		self.hourList[item.TriggerHour] = item
	}
}

func (self *TimerWheelAlarm) Check(t time.Time) {
	fmt.Printf("check at %s\n", t)
	self.mutex.Lock()
	nowSec, nowMin, nowHour := self.curSec, self.curMin, self.curHour
	nowSec += 1
	if nowSec >= 60 {
		nowSec = 0
		nowMin += 1
	}
	if nowMin >= 60 {
		nowMin = 0
		nowHour += 1
	}
	if nowHour >= 24 {
		nowHour = 0
	}
	if nowHour != self.curHour {
		todo := self.hourList[nowHour]
		for {
			if todo == nil {
				break
			}
			item := todo
			todo = todo.next
			self.addToList(item, Minute)
		}
		self.hourList[nowHour] = nil
	}
	if nowMin != self.curMin {
		todo := self.minList[nowMin]
		for {
			if todo == nil {
				break
			}
			item := todo
			todo = todo.next
			self.addToList(item, Second)
		}
		self.minList[nowMin] = nil
	}
	if nowSec != self.curSec {
		todo := self.secList[nowSec]
		for {
			if todo == nil {
				break
			}
			item := todo
			todo = todo.next
			alarm := item.item
			now := time.Now()
			if alarm.When.Before(now.Add(-Interval)) {
				panic(fmt.Sprintf("%s run at wrong time %s", alarm.toStr(), now))
			} else {
				go alarm.run()
			}
		}
		self.secList[nowSec] = nil
	}
	self.curSec, self.curMin, self.curHour = nowSec, nowMin, nowHour
	self.mutex.Unlock()
}

func (self *TimerWheelAlarm) Add(when time.Time, call AlarmFunc, name string) error {
	nowSec, nowMin, nowHour := self.curSec, self.curMin, self.curHour
	item, err := newAlarmListNode(when, call, name)
	if err != nil {
		return err
	}

	item.TriggerSec += nowSec
	item.TriggerMin += nowMin
	item.TriggerHour += nowHour
	if item.TriggerSec >= 60 {
		item.TriggerSec -= 60
		item.TriggerMin += 1
	}
	if item.TriggerMin >= 60 {
		item.TriggerMin -= 60
		item.TriggerHour += 1
	}
	if item.TriggerHour >= 24 {
		return errors.New(fmt.Sprintf("error for add %s : alarm exceed time range", item.item.toStr()))
	}

	self.mutex.Lock()
	if item.TriggerHour != nowHour {
		self.addToList(item, Hour)
	} else if item.TriggerMin != nowMin {
		self.addToList(item, Minute)
	} else if item.TriggerSec != nowSec {
		self.addToList(item, Second)
	}
	self.mutex.Unlock()
	return nil
}

func (self *TimerWheelAlarm) Cancel(name string) error {
	// todo
	return nil
}
