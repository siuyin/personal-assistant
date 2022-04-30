// Package command provides command and control functions.
package command

import (
	log "github.com/sirupsen/logrus"
	"github.com/siuyin/personal-assistant/internal/evt"
)

func init() {
	event("CommandCentreStarted", "command centre started")
	checkForThreats()
}

func checkForThreats() {
	go func() {
		for {
			sub := evt.Sub("pa.>", "commandClient")
			msg := evt.Fetch(sub)
			if msg == "ThreatDetected" {
				event("PoliceHelpRequested", "police help requested")
			}
		}
	}()
}

func info(x ...interface{}) {
	log.WithFields(log.Fields{"module": "command"}).Info(x...)
}
func warn(x ...interface{}) {
	log.WithFields(log.Fields{"module": "command"}).Warn(x...)
}
func debug(x ...interface{}) {
	log.WithFields(log.Fields{"module": "command"}).Debug(x...)
}
func event(name string, msg ...interface{}) {
	evt.Pub("pa.command", name)
	info(msg...)
}
