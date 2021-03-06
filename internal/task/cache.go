package task

import (
	"github.com/Feoks/echo/pkg/task"
	"sync"
)

type cache struct {
	sync.RWMutex

	tasks    []*task.Task
	tasksMap map[uint64]*task.Task
}

func initCache() Cache {
	return &cache{
		tasks:    make([]*task.Task, 0),
		tasksMap: make(map[uint64]*task.Task, 0),
	}
}

func (c *cache) Add(task *task.Task) error {
	c.RWMutex.Lock()
	defer c.RWMutex.Unlock()

	c.tasks = append(c.tasks, task)
	c.tasksMap[task.Id] = task

	return nil
}

func (c *cache) Get(id uint64) (*task.Task, error) {
	c.RWMutex.RLock()
	defer c.RWMutex.RUnlock()

	return c.tasks[id], nil
}

func (c *cache) GetAll() ([]*task.Task, error) {
	c.RWMutex.RLock()
	defer c.RWMutex.RUnlock()

	return c.tasks, nil
}

func (c *cache) Update(task *task.Task) error {
	c.RWMutex.Lock()
	defer c.RWMutex.Unlock()

	c.tasks[task.Id] = task

	return nil
}

func (c *cache) Delete(id uint64) error {
	c.RWMutex.Lock()
	defer c.RWMutex.Unlock()

	delete(c.tasksMap, id)

	return nil
}

func (c *cache) Reset() error {
	c.RWMutex.Lock()
	defer c.RWMutex.Unlock()

	c.tasksMap = make(map[uint64]*task.Task, 0)

	return nil
}

func (c *cache) AddBatch(list []*task.Task) error {
	c.RWMutex.Lock()
	defer c.RWMutex.Unlock()

	for _, t := range list {
		if err := c.Add(t); err != nil {
			return err
		}
	}

	return nil
}
