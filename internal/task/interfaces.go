package task

import "context"

type Repository interface {
	Add(ctx context.Context, task *Task) error
	Get(ctx context.Context, id uint64) (*Task, error)
	Update(ctx context.Context, task *Task) error
	Delete(ctx context.Context, id uint64) error
}

type Cache interface {
	Repository

	Reset() error
}
