package events

import (
	"bytes"
	"context"
	"encoding/gob"
	"encoding/json"
	"github.com/nats-io/nats.go"
	"log"
	"microConsultaPuntos/dominio/entidades"
)

type NatsEventStore struct {
	conn            *nats.Conn
	feedCreatedSub  *nats.Subscription
	feedCreatedChan chan entidades.Usuario
}

func NewNats(url, user, password string) (*NatsEventStore, error) {
	conn, err := nats.Connect(url, nats.UserInfo(user, password))
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
	base, err = json.Marshal(m)
	if err != nil {
		return
	}
	return
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

func (n *NatsEventStore) decodeMessage(data []byte, m interface{}) error {
	b := bytes.Buffer{}
	b.Write(data)
	return gob.NewDecoder(&b).Decode(m)
}

func (n *NatsEventStore) OnCreatedFeed(f func(*entidades.Usuario)) (err error) {
	var topic = "topic_punt_users"
	msg := entidades.Usuario{}
	n.feedCreatedSub, err = n.conn.Subscribe(topic, func(m *nats.Msg) {
		err := n.decodeMessage(m.Data, &msg)
		if err != nil {
			return
		}
		f(&msg)
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