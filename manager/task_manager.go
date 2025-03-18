package manager

import (
	"fmt"
	"reflect"
	"time"

	"task_manager_project/tasks"
)

type TaskManager struct {
	Tasks       []tasks.Task
	FailedTasks []tasks.Task
}

func (tm *TaskManager) AddTask(t tasks.Task) {
	tm.Tasks = append(tm.Tasks, t)
}

func (tm *TaskManager) ExecuteAll() {
	for _, task := range tm.Tasks {
		startTime := time.Now()
		taskType := reflect.TypeOf(task).Elem().Name()

		fmt.Printf("[TaskManager] Running Task: %s | Info: %s\n", taskType, task.Info())

		err := task.Execute()
		duration := time.Since(startTime)

		if err != nil {
			fmt.Printf("[TaskManager] ❌ Task %s FAILED: %s (Duration: %v)\n", taskType, err, duration)
			tm.FailedTasks = append(tm.FailedTasks, task)
		} else {
			fmt.Printf("[TaskManager] ✅ Task %s SUCCESS (Duration: %v)\n", taskType, duration)
		}
	}
}

func (tm *TaskManager) RetryFailedTasks() {
	if len(tm.FailedTasks) == 0 {
		fmt.Println("[TaskManager] No failed tasks to retry.")
		return
	}

	fmt.Println("[TaskManager] Retrying failed tasks...")
	tasksToRetry := tm.FailedTasks
	tm.FailedTasks = nil

	for _, task := range tasksToRetry {
		startTime := time.Now()
		taskType := reflect.TypeOf(task).Elem().Name()

		fmt.Printf("[TaskManager] Retrying Task: %s | Info: %s\n", taskType, task.Info())

		err := task.Execute()
		duration := time.Since(startTime)

		if err != nil {
			fmt.Printf("[TaskManager] ❌ RETRY Task %s FAILED: %s (Duration: %v)\n", taskType, err, duration)
			tm.FailedTasks = append(tm.FailedTasks, task)
		} else {
			fmt.Printf("[TaskManager] ✅ RETRY Task %s SUCCESS (Duration: %v)\n", taskType, duration)
		}
	}
}
