package main

import (
	"math/rand"
	"time"

	"task_manager_project/manager"
	"task_manager_project/tasks"
)

func main() {
	rand.Seed(time.Now().UnixNano())

	manager := manager.TaskManager{}

	manager.AddTask(&tasks.EmailTask{BaseTask: tasks.NewBaseTask("1", "Send welcome email")})
	manager.AddTask(&tasks.SMSTask{BaseTask: tasks.NewBaseTask("2", "Send OTP SMS")})
	manager.AddTask(&tasks.ReportTask{BaseTask: tasks.NewBaseTask("3", "Generate monthly report")})

	manager.ExecuteAll()

	manager.RetryFailedTasks()
}
