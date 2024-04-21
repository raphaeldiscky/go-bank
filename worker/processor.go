package worker

import (
	"context"

	"github.com/hibiken/asynq"
	db "github.com/raphaeldiscky/simple-bank/db/sqlc"
)

const (
	QueueCritical = "critical"
	QueueDefault  = "default"
)

type TaskProcessor interface {
	Start() error
	ProcessTaskSendVerifyEmail(ctx context.Context, task *asynq.Task) error
}

type RedisTaskProcessor struct {
	server *asynq.Server
	store  db.Store
}

// NewRedisTaskProcessor will pick up the tasks from the Redis queue and process them
func NewRedisTaskProcessor(redisOpt asynq.RedisClientOpt, store db.Store) TaskProcessor {
	server := asynq.NewServer(
		redisOpt,
		asynq.Config{
			// Specify how many workers you want
			Concurrency: 10,
			// Optionally specify multiple queues with different priority.
			// If not specified, uses the default (10).
			Queues: map[string]int{
				QueueCritical: 6,
				QueueDefault:  3,
			},
		},
	)

	return &RedisTaskProcessor{
		server: server,
		store:  store,
	}
}

func (processor *RedisTaskProcessor) Start() error {
	mux := asynq.NewServeMux()

	mux.HandleFunc(TaskSendVerifyEmail, processor.ProcessTaskSendVerifyEmail)
	return processor.server.Start(mux)
}
