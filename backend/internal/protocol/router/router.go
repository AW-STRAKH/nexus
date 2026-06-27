package router

import (
	"context"
	"fmt"
	"sync"

	nexusproto "github.com/awatansh/nexus/internal/protocol/proto"
)

type Router struct {
	mu       sync.RWMutex
	handlers map[nexusproto.MessageType]Handler
}

func NewRouter() *Router {

	return &Router{
		handlers: make(
			map[nexusproto.MessageType]Handler,
		),
	}
}

func (r *Router) Register(
	messageType nexusproto.MessageType,
	handler Handler,
) {

	r.mu.Lock()
	defer r.mu.Unlock()

	r.handlers[messageType] = handler
}

func (r *Router) Dispatch(
	ctx context.Context,
	envelope *nexusproto.Envelope,
) error {

	r.mu.RLock()
	handler, found :=
		r.handlers[envelope.Type]
	r.mu.RUnlock()

	if !found {
		return fmt.Errorf(
			"no handler registered for %v",
			envelope.Type,
		)
	}

	return handler.Handle(
		ctx,
		envelope,
	)
}
