package tasks

type Task interface {
	Execute() error
	Info() string
}
