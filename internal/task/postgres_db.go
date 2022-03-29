package task

import (
	"context"
	"errors"
	"git.repo.services.lenvendo.ru/grade-factor/echo/internal/db/postgres"
)

const (
	AddTaskQuery    = `insert into "task"."tasks" ("id", "name", "active") values($1, $2, $3)`
	GetTaskQuery    = `select * from "task"."tasks" where id=$1`
	UpdateTaskQuery = `update "task"."tasks" set name=$2, active=$3 where id=$1`
	DeleteTaskQuery = `delete from "task"."tasks" where id=$1`
	GetAllQuery     = `select * from "task"."tasks"`
)

type postgresDb struct {
	db  postgres.Connection
	ctx context.Context
}

func newPostgresConnection(conn postgres.Connection, ctx context.Context) Repository {
	return &postgresDb{
		db:  conn,
		ctx: ctx,
	}
}

func (p *postgresDb) execMasterQuery(query string, args ...interface{}) error {
	conn, err := p.db.GetMasterConn(p.ctx)
	if err != nil {
		return err
	}
	defer conn.Release()

	tx, err := conn.Begin(p.ctx)
	if err != nil {
		return err
	}

	result, err := tx.Exec(p.ctx, query, args...)
	if err != nil {
		return err
	}
	if result.RowsAffected() < 1 {
		return errors.New("no row processed")
	}

	err = tx.Commit(p.ctx)
	if err != nil {
		return err
	}

	return nil
}

func (p *postgresDb) Add(task *Task) error {
	return p.execMasterQuery(
		AddTaskQuery,
		task.Id,
		task.Name,
		task.Active,
	)
}

func (p *postgresDb) Delete(id uint64) error {
	return p.execMasterQuery(
		DeleteTaskQuery,
		id,
	)
}

func (p *postgresDb) Update(task *Task) error {
	return p.execMasterQuery(
		UpdateTaskQuery,
		task.Id,
		task.Name,
		task.Active,
	)
}

func (p *postgresDb) Get(id uint64) (*Task, error) {
	conn, err := p.db.GetReplicaConn(p.ctx)
	if err != nil {
		return nil, err
	}
	defer conn.Release()

	t := Task{}
	if err := conn.QueryRow(
		p.ctx,
		GetTaskQuery,
		id,
	).Scan(&t); err != nil {
		return nil, err
	}

	return &t, nil
}

func (p *postgresDb) GetAll() ([]*Task, error) {
	conn, err := p.db.GetReplicaConn(p.ctx)
	if err != nil {
		return nil, err
	}
	defer conn.Release()

	rows, err := conn.Query(p.ctx, GetAllQuery)
	if err != nil {
		return nil, err
	}

	list := make([]*Task, 0)
	for rows.Next() {
		t := Task{}
		if err = rows.Scan(&t); err != nil {
			return nil, err
		}

		list = append(list, &t)
	}

	return list, nil
}
