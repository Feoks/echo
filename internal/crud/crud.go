package crud

import "git.repo.services.lenvendo.ru/grade-factor/echo/internal/db/postgres"

type crud struct {
	itemById map[uint32]*Crud
	all      []*Crud

	db postgres.Connection
}

type Crud struct {
	id       uint32
	name     string
	required bool
}

type ICrud interface {
	Create() (*Crud, error)
	Read() ([]*Crud, error)
	Update() (*Crud, error)
	Delete() error

	Reset()
}

func NewCrud(db postgres.Connection) ICrud {
	return &crud{
		itemById: make(map[uint32]*Crud, 0),
		all:      make([]*Crud, 0, 0),
		db:       db,
	}
}

func (c *crud) Create() (*Crud, error) {
	return nil, nil
}

func (c *crud) Read() ([]*Crud, error) {
	return c.all, nil
}

func (c *crud) Update() (*Crud, error) {
	//TODO
}

func (c *crud) Delete() error {
	return nil
}

func (c *crud) Reset() {
	//TODO
}
