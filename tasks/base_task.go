package tasks

import "time"

type BaseTask struct {
	ID          string
	CreatedAt   time.Time
	Description string
}

func NewBaseTask(id, desc string) BaseTask {
	return BaseTask{
		ID:          id,
		CreatedAt:   time.Now(),
		Description: desc,
	}
}
