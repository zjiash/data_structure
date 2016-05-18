package alarm

import (
	"fmt"
	"testing"
	"time"
)

type task struct {
	Name string
	When time.Time
}

func (self *task) run() {
	fmt.Println("#################")
	fmt.Println("task running...")
	fmt.Println("task done...")
	fmt.Println("#################")
}

func warpTask(t task) func() {
	return func() {
		t.run()
	}
}

func TestHeapAlarm(test *testing.T) {
	now := time.Now()
	alarm := NewHeapAlarm()

	task0 := task{
		Name: "task0",
		When: now.Add(time.Second * 2),
	}
	task1 := task{
		Name: "task1",
		When: now.Add(time.Second * 8),
	}
	task2 := task{
		Name: "task2",
		When: now.Add(time.Second * 12),
	}
	alarm.Add(task0.When, warpTask(task0), task0.Name)
	time.Sleep(4 * time.Second)

	alarm.Add(task1.When, warpTask(task1), task1.Name)
	alarm.Add(task2.When, warpTask(task2), task2.Name)
	alarm.Add(task1.When, warpTask(task1), task1.Name)

	time.Sleep(12 * time.Second)
}
