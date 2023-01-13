package eventos

import (
	"context"
	"microMutationPuntos/dominio/entidades"
)

type EventStore interface {
	Close()
	PublishCreated(ctx context.Context, feed *entidades.Usuario, topic string) error
	SubscribeCreated(ctx context.Context, topic string) (<-chan entidades.Usuario, error)
	OnCreated(f func(entidades.Usuario), topic string) error
}

var eventStore EventStore

func SetEventStore(store EventStore) {
	eventStore = store
}

func Close() {
	eventStore.Close()
}

func PublishCreated(ctx context.Context, feed *entidades.Usuario, topic string) error {
	return eventStore.PublishCreated(ctx, feed, topic)
}

func SubscribeCreated(ctx context.Context, topic string) (<-chan entidades.Usuario, error) {
	return eventStore.SubscribeCreated(ctx, topic)
}

func OnCreated(ctx context.Context, f func(entidades.Usuario), topic string) error {
	return eventStore.OnCreated(f, topic)
}
