package step

import (
	"context"
)

type Todo struct {
	// TODO 步骤所需字段
}

func NewTodo() *Todo {
	return &Todo{
		// 字段初步化
	}
}

func (t *Todo) Runnable() bool {
	return true
}

func (t *Todo) Run(_ context.Context) (err error) {
	return
}
