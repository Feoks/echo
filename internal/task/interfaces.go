package task

import "github.com/Feoks/echo/pkg/task"

type Repository interface {
	Add(task *task.Task) error
	Get(id uint64) (*task.Task, error)
	GetAll() ([]*task.Task, error)
	Update(task *task.Task) error
	Delete(id uint64) error
}

type Cache interface {
	Repository

	Reset() error
	AddBatch(list []*task.Task) error
}
