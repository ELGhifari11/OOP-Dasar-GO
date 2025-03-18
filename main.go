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

	// Menambahkan tugas
	manager.AddTask(&tasks.EmailTask{BaseTask: tasks.NewBaseTask("1", "Send welcome email")})
	manager.AddTask(&tasks.SMSTask{BaseTask: tasks.NewBaseTask("2", "Send OTP SMS")})
	manager.AddTask(&tasks.ReportTask{BaseTask: tasks.NewBaseTask("3", "Generate monthly report")})

	// Jalankan semua tugas
	manager.ExecuteAll()

	// Retry jika ada tugas gagal
	manager.RetryFailedTasks()
}
