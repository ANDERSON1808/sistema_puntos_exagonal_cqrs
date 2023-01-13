package eventos

import (
	"context"
	"microConsultaPuntos/dominio/entidades"
)

type EventStore interface {
	Close()
	PublishCreated(ctx context.Context, feed *entidades.Usuario, topic string) error
	SubscribeCreated(ctx context.Context, topic string) (<-chan entidades.Usuario, error)
	OnCreatedFeed(f func(entidades.Usuario)) (err error)
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

func OnCreatedFeed(f func(entidades.Usuario)) (err error) {
	return eventStore.OnCreatedFeed(f)
}
