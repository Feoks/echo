package task

import "context"

type cache struct {
	tasks map[uint64]*Task
}

func initCache() Cache {
	return &cache{
		tasks: make(map[uint64]*Task, 0),
	}
}

func (c *cache) Add(ctx context.Context, task *Task) error {
	c.tasks[task.Id] = task

	return nil
}

func (c *cache) Get(ctx context.Context, id uint64) (*Task, error) {
	//if c.tasks[id] == nil

	return c.tasks[id], nil
}

func (c *cache) Update(ctx context.Context, task *Task) error {
	c.tasks[task.Id] = task

	return nil
}

func (c *cache) Delete(ctx context.Context, id uint64) error {
	delete(c.tasks, id)

	return nil
}

func (c *cache) Reset() error {
	c.tasks = make(map[uint64]*Task, 0)

	return nil
}
