package tasks

import (
	"fmt"
	"math/rand"
	"time"
)

type EmailTask struct {
	BaseTask
}

func (e *EmailTask) Execute() error {
	fmt.Printf("[EmailTask] Executing: %s\n", e.Description)
	time.Sleep(time.Duration(rand.Intn(3)) * time.Second)
	if rand.Intn(5) == 0 {
		return fmt.Errorf("[EmailTask] Failed to send email!")
	}
	fmt.Printf("[EmailTask] Completed: %s\n", e.Description)
	return nil
}

func (e *EmailTask) Info() string {
	return fmt.Sprintf("Task: Email | ID: %s | Desc: %s", e.ID, e.Description)
}
