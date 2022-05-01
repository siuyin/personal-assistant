// Package stor abstracts persistence and storage.
package stor

import (
	log "github.com/sirupsen/logrus"
	"github.com/siuyin/personal-assistant/internal/evt"
)

func init() {
	debug("storage system started")
}

func info(x ...interface{}) {
	log.WithFields(log.Fields{"module": "stor"}).Info(x...)
}
func warn(x ...interface{}) {
	log.WithFields(log.Fields{"module": "stor"}).Warn(x...)
}
func debug(x ...interface{}) {
	log.WithFields(log.Fields{"module": "stor"}).Debug(x...)
}
func event(name string, msg ...interface{}) {
	evt.Pub("pa.stor", name)
	info(msg...)
}
