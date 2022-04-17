package task

import (
	"context"
	"github.com/Feoks/echo/internal/db/postgres"
	"github.com/Feoks/echo/pkg/task"
)

type defaultRepository struct {
	postgres Repository
	cache    Cache
}

func NewRepository(conn postgres.Connection, ctx context.Context) Repository {
	return &defaultRepository{
		postgres: newPostgresConnection(conn, ctx),
		cache:    initCache(),
	}
}

func (r *defaultRepository) Add(task *task.Task) error {
	if err := r.postgres.Add(task); err != nil {
		return err
	}

	if err := r.cache.Add(task); err != nil {
		return err
	}

	return nil
}

func (r *defaultRepository) Get(id uint64) (*task.Task, error) {
	task, err := r.cache.Get(id)
	if err != nil {
		return nil, err
	}
	if task != nil {
		return task, nil
	}

	task, err = r.postgres.Get(id)
	if err != nil {
		return nil, err
	}
	if task == nil {
		return nil, nil
	}

	err = r.cache.Add(task)
	if err != nil {
		return nil, err
	}

	return task, nil
}

func (r *defaultRepository) Delete(id uint64) error {
	if err := r.postgres.Delete(id); err != nil {
		return err
	}

	if err := r.cache.Delete(id); err != nil {
		return err
	}

	return nil
}

func (r *defaultRepository) Update(task *task.Task) error {
	if err := r.postgres.Update(task); err != nil {
		return err
	}

	if err := r.cache.Update(task); err != nil {
		return err
	}

	return nil
}

func (r *defaultRepository) GetAll() ([]*task.Task, error) {
	list, err := r.cache.GetAll()
	if err == nil {
		return list, nil
	}

	list, err = r.postgres.GetAll()
	if err != nil {
		return nil, err
	}

	if err := r.cache.AddBatch(list); err != nil {
		return nil, err
	}

	return list, nil
}
