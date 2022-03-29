package task

type Repository interface {
	Add(task *Task) error
	Get(id uint64) (*Task, error)
	GetAll() ([]*Task, error)
	Update(task *Task) error
	Delete(id uint64) error
}

type Cache interface {
	Repository

	Reset() error
	AddBatch(list []*Task) error
}
