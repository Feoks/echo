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
)

type postgresDb struct {
	db postgres.Connection
}

func newPostgresConnection(conn postgres.Connection) Repository {
	return &postgresDb{
		db: conn,
	}
}

func (p *postgresDb) execMasterQuery(ctx context.Context, query string, args ...interface{}) err {
	conn, err := p.db.GetMasterConn(ctx)
	if err != nil {
		return err
	}
	defer conn.Release()

	tx, err := conn.Begin(ctx)
	if err != nil {
		return err
	}

	result, err := tx.Exec(ctx, query, args...)
	if err != nil {
		return err
	}
	if result.RowsAffected() < 1 {
		return errors.New("no row processed")
	}

	err = tx.Commit(ctx)
	if err != nil {
		return err
	}
}

func (p *postgresDb) Add(ctx context.Context, task *Task) error {
	p.execMasterQuery(
		ctx,
		AddTaskQuery,
		task.Id,
		task.Name,
		task.Active,
	)

	return nil
}

func (p *postgresDb) Get(ctx context.Context, id uint64) (*Task, error) {
	conn, err := p.db.GetReplicaConn(ctx)
	if err != nil {
		conn, err = p.db.GetMasterConn(ctx)
		if err != nil {
			return nil, err
		}
	}
	defer conn.Release()

	row := conn.QueryRow(
		ctx,
		GetTaskQuery,
		id,
	)

	if row == nil {
		return nil, nil
	}

	var task *Task

	if err = row.Scan(&task); err != nil {
		return nil, err
	}

	return task, nil
}

func (p *postgresDb) Delete(ctx context.Context, id uint64) error {
	p.execMasterQuery(
		ctx,
		DeleteTaskQuery,
		task.Id,
	)

	return nil
}

func (p *postgresDb) Update(ctx context.Context, task *Task) error {
	p.execMasterQuery(
		ctx,
		UpdateTaskQuery,
		task.Id,
		task.Name,
		task.Active,
	)

	return nil
}
