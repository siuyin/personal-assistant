package outp

import (
	log "github.com/sirupsen/logrus"
	"github.com/siuyin/personal-assistant/internal/evt"
)

func Init() {
	debug("output system started")

	checkForCommands()
}
func checkForCommands() {
	go func() {
		for {
			sub := evt.Sub("pa.>", "outpClient")
			msg := evt.Fetch(sub)
			if msg == "PoliceHelpRequested" {
				execPoliceHelpCmd()
			}
		}
	}()
}

func execPoliceHelpCmd() {
	debug("location requested from learning module")
	event("PoliceHelpRequestCmdFormatted", "police help request command formatted")
	event("PoliceHelpRequestSent", "police help request sent")
}

func info(x ...interface{}) {
	log.WithFields(log.Fields{"module": "outp"}).Info(x...)
}
func warn(x ...interface{}) {
	log.WithFields(log.Fields{"module": "outp"}).Warn(x...)
}
func debug(x ...interface{}) {
	log.WithFields(log.Fields{"module": "outp"}).Debug(x...)
}
func event(name string, msg ...interface{}) {
	evt.Pub("pa.outp", name)
	info(msg...)
}
