package tasks

import (
	"fmt"
	"math/rand"
	"time"
)

type SMSTask struct {
	BaseTask
}

func (s *SMSTask) Execute() error {
	fmt.Printf("[SMSTask] Executing: %s\n", s.Description)
	time.Sleep(time.Duration(rand.Intn(3)) * time.Second)
	if rand.Intn(5) == 0 {
		return fmt.Errorf("[SMSTask] Failed to send SMS!")
	}
	fmt.Printf("[SMSTask] Completed: %s\n", s.Description)
	return nil
}

func (s *SMSTask) Info() string {
	return fmt.Sprintf("Task: SMS | ID: %s | Desc: %s", s.ID, s.Description)
}
