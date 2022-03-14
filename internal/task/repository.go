package task

import (
	"context"
	"git.repo.services.lenvendo.ru/grade-factor/echo/internal/db/postgres"
)

type defaultRepository struct {
	postgres Repository
	cache    Repository
}

func NewRepository(conn postgres.Connection) Repository {
	return &defaultRepository{
		postgres: newPostgresConnection(conn),
		cache:    initCache(),
	}
}

func (r *defaultRepository) Add(ctx context.Context, task *Task) error {
	if err := r.postgres.Add(ctx, task); err != nil {
		return err
	}

	if err := r.cache.Add(ctx, task); err != nil {
		return err
	}

	return nil
}

func (r *defaultRepository) Get(ctx context.Context, id uint64) (*Task, error) {
	task, err := r.cache.Get(ctx, id)
	if err != nil {
		return nil, err
	}
	if task != nil {
		return task, nil
	}

	task, err = r.postgres.Get(ctx, id)
	if err != nil {
		return nil, err
	}
	if task == nil {
		return nil, nil
	}

	err = r.cache.Add(ctx, task)
	if err != nil {
		return nil, err
	}

	return task, nil
}

func (r *defaultRepository) Delete(ctx context.Context, id uint64) error {
	if err := r.postgres.Delete(ctx, id); err != nil {
		return err
	}

	if err := r.cache.Delete(ctx, id); err != nil {
		return err
	}

	return nil
}

func (r *defaultRepository) Update(ctx context.Context, task *Task) error {
	err = r.postgres.Update(ctx, id)
	if err != nil {
		return err
	}

	err = r.cache.Update(ctx, id)
	if err != nil {
		return err
	}

	return nil
}
