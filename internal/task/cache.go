package task

import (
	"context"
	"sync"
)

type cache struct {
	sync.RWMutex

	tasks    []*Task
	tasksMap map[uint64]*Task
}

func initCache() Cache {
	return &cache{
		tasks:    make([]*Task, 0),
		tasksMap: make(map[uint64]*Task, 0),
	}
}

func (c *cache) Add(ctx context.Context, task *Task) error {
	c.RWMutex.Lock()
	defer c.RWMutex.Unlock()

	c.tasks = append(c.tasks, task)
	c.tasksMap[task.Id] = task

	return nil
}

func (c *cache) Get(ctx context.Context, id uint64) (*Task, error) {
	c.RWMutex.RLock()
	defer c.RWMutex.RUnlock()

	return c.tasks[id], nil
}

func (c *cache) GetAll(ctx context.Context) ([]*Task, error) {
	c.RWMutex.RLock()
	defer c.RWMutex.RUnlock()

	return c.tasks, nil
}

func (c *cache) Update(ctx context.Context, task *Task) error {
	c.RWMutex.Lock()
	defer c.RWMutex.Unlock()

	c.tasks[task.Id] = task

	return nil
}

func (c *cache) Delete(ctx context.Context, id uint64) error {
	c.RWMutex.Lock()
	defer c.RWMutex.Unlock()

	delete(c.tasksMap, id)

	return nil
}

func (c *cache) Reset() error {
	c.RWMutex.Lock()
	defer c.RWMutex.Unlock()

	c.tasksMap = make(map[uint64]*Task, 0)

	return nil
}
