package tasks

import (
	"fmt"
	"math/rand"
	"time"
)

type ReportTask struct {
	BaseTask
}

func (r *ReportTask) Execute() error {
	fmt.Printf("[ReportTask] Executing: %s\n", r.Description)
	time.Sleep(time.Duration(rand.Intn(3)) * time.Second)
	if rand.Intn(5) == 0 {
		return fmt.Errorf("[ReportTask] Failed to generate report!")
	}
	fmt.Printf("[ReportTask] Completed: %s\n", r.Description)
	return nil
}

func (r *ReportTask) Info() string {
	return fmt.Sprintf("Task: Report | ID: %s | Desc: %s", r.ID, r.Description)
}
