// Package evt provides event management functions.
package evt

import (
	log "github.com/sirupsen/logrus"

	"github.com/nats-io/nats.go"
	"github.com/siuyin/dflt"
)

// Pub publishes an event.
func Pub(subj, s string) {
	if _, err := js.Publish(subj, []byte(s)); err != nil {
		warn("evt: pub: ", err)
	}
	debug("published to ", subj, ": ", s)

}
func warn(x ...interface{}) {
	log.WithFields(log.Fields{"module": "evt"}).Warn(x...)
}
func info(x ...interface{}) {
	log.WithFields(log.Fields{"module": "evt"}).Info(x...)
}
func debug(x ...interface{}) {
	log.WithFields(log.Fields{"module": "evt"}).Debug(x...)
}
func fatal(x ...interface{}) {
	log.WithFields(log.Fields{"module": "evt", "src": "internal/evt/evt.go"}).Fatal(x...)
}

// Sub subscribes to a topic / subject.
func Sub(subj, clientName string) *nats.Subscription {
	sub, err := js.PullSubscribe(subj, clientName)
	if err != nil {
		warn("sub: ", err)
	}
	return sub
}

// Fetch fetches an event from a subscription and returns the event.
func Fetch(sub *nats.Subscription) string {
	ms, err := sub.Fetch(1)
	if err != nil && err.Error() == "nats: timeout" {
		return ""
	}
	if err != nil {
		warn("fetch: ", err)
		return ""
	}

	if err := ms[0].Ack(); err != nil {
		warn("fetch ack:", err)
	}
	dat := string(ms[0].Data)
	debug("fetched: ", dat)
	return dat
}

func init() {
	if lvl := dflt.EnvString("LogLevel", "info"); lvl != "info" {
		log.SetLevel(log.DebugLevel)
	}
	initNATS()
	debug("event management system initialised")
}

var (
	nc  *nats.Conn
	js  nats.JetStreamContext
	err error
)

func initNATS() {
	nc, err = nats.Connect(dflt.EnvString("NATS_Svr", "nats://acc:acc@localhost:4223"))
	if err != nil {
		fatal("connect: ", err)
	}

	js, err = nc.JetStream(nats.PublishAsyncMaxPending(256))
	if err != nil {
		fatal("jetstream: ", err)
	}

	setupNATSStreams()
}

func setupNATSStreams() {
	if _, err := js.AddStream(&nats.StreamConfig{
		Name:     "pa", // personal assistant
		Subjects: []string{"pa.>"},
	}); err != nil {
		fatal("setupNATSStreams: ", err)
	}
}
