package alarm

import (
	"fmt"
	"time"
)

const MilliSecondsPerTicker int = 1000
const Interval time.Duration = time.Millisecond * time.Duration(MilliSecondsPerTicker)

type AlarmFunc func()

type Alarm interface {
	Add(when time.Time, call AlarmFunc, name string) error
	Cancel(name string) error
}

type alarmItem struct {
	When time.Time
	Call AlarmFunc
	Name string
}

func (self *alarmItem) Priority() int64 {
	return -self.When.Unix()
}

func (self *alarmItem) toStr() string {
	return fmt.Sprintf("alarm with name = %s, when = %s", self.Name, self.When)
}

func (self *alarmItem) run() {
	fmt.Printf("%s start run in %s\n", self.toStr(), time.Now())
	self.Call()
}

func newAlarmItem(when time.Time, call AlarmFunc, name string) *alarmItem {
	return &alarmItem{
		When: when,
		Call: call,
		Name: name,
	}
}
