// Package evt provides event management functions.
package evt

import (
	"log"

	"github.com/nats-io/nats.go"
	"github.com/siuyin/dflt"
)

// Pub published an event.
func Pub(s string) {
	if _, err := js.Publish("pa", []byte(s)); err != nil {
		log.Println("evt: pub: ", err)
	}
}

// Sub subscribes to a topic / subject.
func Sub(clientName string) *nats.Subscription {
	if _, err := js.AddConsumer("pa", &nats.ConsumerConfig{
		Durable:       clientName,
		FilterSubject: "pa",
		AckPolicy:     nats.AckExplicitPolicy,
	}); err != nil {
		log.Println("evt: sub: addconsumer: ", err)
	}

	sub, err := js.PullSubscribe("pa", clientName)
	if err != nil {
		log.Println("evt: sub: pullsubscribe: ", err)
	}
	return sub
}

func logST(s string) { // log skip in test mode
	mod := dflt.EnvString("Mode", "test")
	if mod == "test" {
		return
	}
	log.Println(s)
}

func init() {
	initNATS()
	logST("evt: event management system initialised")
}

var (
	nc  *nats.Conn
	js  nats.JetStreamContext
	err error
)

func initNATS() {
	nc, err = nats.Connect(dflt.EnvString("NATS_Svr", "localhost:4223"))
	if err != nil {
		log.Fatal("evt: connect: ", err)
	}

	js, err = nc.JetStream(nats.PublishAsyncMaxPending(256))
	if err != nil {
		log.Fatal("evt: jetstream: ", err)
	}

	setupNATSStreams()
}

func setupNATSStreams() {
	if _, err := js.AddStream(&nats.StreamConfig{
		Name:     "pa", // personal assistant
		Subjects: []string{"pa"},
	}); err != nil {
		log.Fatal("evt: setupNATSStreams: ", err)
	}
}
