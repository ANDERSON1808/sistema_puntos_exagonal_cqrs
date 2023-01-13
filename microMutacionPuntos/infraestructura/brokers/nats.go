package events

import (
	"bytes"
	"context"
	"encoding/gob"
	"github.com/nats-io/nats.go"
	"log"
	"microMutationPuntos/dominio/entidades"
	"microMutationPuntos/infraestructura/Modelos"
)

type NatsEventStore struct {
	conn            *nats.Conn
	feedCreatedSub  *nats.Subscription
	feedCreatedChan chan entidades.Usuario
}

func NewNats(url string) (*NatsEventStore, error) {
	conn, err := nats.Connect(url)
	if err != nil {
		return nil, err
	}
	log.Printf("Star server nats... %s", url)
	return &NatsEventStore{
		conn: conn,
	}, nil
}

func (n *NatsEventStore) Close() {
	if n.conn != nil {
		n.conn.Close()
	}
	if n.feedCreatedSub != nil {
		err := n.feedCreatedSub.Unsubscribe()
		if err != nil {
			return
		}
	}
	close(n.feedCreatedChan)
}

func (n *NatsEventStore) encodeMessage(m *entidades.Usuario) (base []byte, err error) {
	b := bytes.Buffer{}
	err = gob.NewEncoder(&b).Encode(m)
	if err != nil {
		return nil, err
	}
	return b.Bytes(), nil
}

func (n *NatsEventStore) PublishCreated(_ context.Context, msg *entidades.Usuario, topic string) error {
	log.Printf("Start process publish nats, topic: %v, msn: %x\n", topic, msg)
	data, err := n.encodeMessage(msg)
	if err != nil {
		log.Printf("Error in encoding msn: %e", err)
		return err
	}
	return n.conn.Publish(topic, data)
}
func (n *NatsEventStore) PublishRedimir(ctx context.Context, feed Modelos.RedimirPuntos, topic string) error {
	log.Printf("Start process publish nats, topic: %v, msn: %x\n", topic, feed)
	b := bytes.Buffer{}
	err := gob.NewEncoder(&b).Encode(feed)
	if err != nil {
		return err
	}
	return n.conn.Publish(topic, b.Bytes())
}

func (n *NatsEventStore) decodeMessage(data []byte, m interface{}) error {
	b := bytes.Buffer{}
	b.Write(data)
	return gob.NewDecoder(&b).Decode(m)
}

func (n *NatsEventStore) OnCreated(f func(entidades.Usuario), topic string) (err error) {
	msg := entidades.Usuario{}
	n.feedCreatedSub, err = n.conn.Subscribe(topic, func(m *nats.Msg) {
		err := n.decodeMessage(m.Data, &msg)
		if err != nil {
			return
		}
		f(msg)
	})
	return
}

func (n *NatsEventStore) SubscribeCreated(_ context.Context, topic string) (<-chan entidades.Usuario, error) {
	m := entidades.Usuario{}
	n.feedCreatedChan = make(chan entidades.Usuario, 64)
	ch := make(chan *nats.Msg, 64)
	var err error
	n.feedCreatedSub, err = n.conn.ChanSubscribe(topic, ch)
	if err != nil {
		return nil, err
	}
	go func() {
		for {
			select {
			case msg := <-ch:
				err := n.decodeMessage(msg.Data, &m)
				if err != nil {
					return
				}
				n.feedCreatedChan <- m
			}
		}
	}()
	return (<-chan entidades.Usuario)(n.feedCreatedChan), nil
}
