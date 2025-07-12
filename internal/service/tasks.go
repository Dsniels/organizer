package service

import "context"

type ITask interface {
	ListTasks(ctx context.Context)
	CreateTask(ctx context.Context)
	GetTask(ctx context.Context)
}
