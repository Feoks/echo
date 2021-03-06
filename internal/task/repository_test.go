package task

import (
	"context"
	"fmt"
	"github.com/Feoks/echo/configs"
	"github.com/Feoks/echo/internal/db/postgres"
	"github.com/Feoks/echo/pkg/task"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestNewRepository(t *testing.T) {
	t1 := &task.Task{
		Id:     1,
		Name:   "First",
		Active: true,
	}

	a := assert.New(t)
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()
	cfg := configs.NewConfig()
	if err := cfg.Read(); err != nil {
		fmt.Println("Config error: ")
		fmt.Println(err.Error())
		return
	}

	conn, err := postgres.NewConnection(ctx, cfg)
	if err != nil {
		fmt.Println("Connection error: ")
		fmt.Println(err.Error())
		return
	}

	r := NewRepository(conn, ctx)

	t.Run("Add & Read", func(t *testing.T) {
		err := r.Add(t1)
		a.Nil(err)

		task, err := r.Get(1)
		a.Equal(task, t1)
		a.Nil(err)
	})
}
