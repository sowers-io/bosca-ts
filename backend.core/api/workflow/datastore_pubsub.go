package workflow

import (
	"bosca.io/api/protobuf/bosca/workflow"
	"bosca.io/pkg/configuration"
	"context"
	"encoding/json"
	"github.com/redis/go-redis/v9"
	"sync"
)

type PubSub struct {
	client        *redis.Client
	mutex         sync.Mutex
	subscriptions map[string]*redis.PubSub
}

func NewPubSub(cfg *configuration.ClientEndpoints) *PubSub {
	return &PubSub{
		client: redis.NewClient(&redis.Options{
			Addr: cfg.RedisAddress,
		}),
	}
}

func (ps *PubSub) Publish(ctx context.Context, queue string, notification *workflow.WorkflowExecutionNotification) error {
	data, err := json.Marshal(notification)
	if err != nil {
		return err
	}
	result := ps.client.Publish(ctx, queue, data)
	return result.Err()
}

func (ps *PubSub) Subscribe(ctx context.Context, queue string) {
	ps.mutex.Lock()
	defer ps.mutex.Unlock()
	if ps.subscriptions[queue] != nil {
		return
	}
	ps.subscriptions[queue] = ps.client.Subscribe(ctx, queue)
}

func (ps *PubSub) Listen(ctx context.Context, queue string, processor func(notification *workflow.WorkflowExecutionNotification) (bool, error)) error {
	ps.mutex.Lock()
	subscription := ps.subscriptions[queue]
	if subscription == nil {
		ps.mutex.Unlock()
		ps.Subscribe(ctx, queue)
		ps.mutex.Lock()
		subscription = ps.subscriptions[queue]
	}
	ps.mutex.Unlock()
	for msg := range subscription.Channel() {
		notification := &workflow.WorkflowExecutionNotification{}
		err := json.Unmarshal([]byte(msg.Payload), notification)
		if err != nil {
			return err
		}
		if processNext, err := processor(notification); !processNext || err != nil {
			return err
		}
	}
	return nil
}

func (ps *PubSub) Unsubscribe(queue string) error {
	ps.mutex.Lock()
	defer ps.mutex.Unlock()
	subscription := ps.subscriptions[queue]
	if subscription != nil {
		delete(ps.subscriptions, queue)
		return subscription.Close()
	}
	return nil
}
