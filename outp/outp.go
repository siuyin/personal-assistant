package outp

import (
	log "github.com/sirupsen/logrus"
	"github.com/siuyin/personal-assistant/internal/evt"
)

func init() {
	debug("output system started")
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
