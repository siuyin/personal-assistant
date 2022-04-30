// Package learn provides learning and classification functions.
package learn

import (
	log "github.com/sirupsen/logrus"
	"github.com/siuyin/personal-assistant/internal/evt"
)

func init() {
	event("LearningSysStarted", "learning system started")
	event("LocationFound", "location found: Singapore")

	debug("door open detected")
	debug("person detected")
	debug("weapon (knife) detected")

	event("ThreatDetected", "threat detected")
}

func info(x ...interface{}) {
	log.WithFields(log.Fields{"module": "learn"}).Info(x...)
}
func debug(x ...interface{}) {
	log.WithFields(log.Fields{"module": "learn"}).Debug(x...)
}
func event(name string, msg ...interface{}) {
	evt.Pub("pa.learn", name)
	info(msg...)
}
